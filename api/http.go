package api

import (
	// "io/ioutil"
	// "log"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/jp-chl/test-go-clean-architecture/domain/service"
	js "github.com/jp-chl/test-go-clean-architecture/pkg/serializer"
)

type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService service.RedirectService
}

func NewHandler(redirectService service.RedirectService) RedirectHandler {
	return &handler{redirectService: redirectService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *handler) serializer(contentType string) service.RedirectSerializer {
	return &js.Redirect{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code, err := h.getCodePathParameter(r)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	redirect, err := h.redirectService.Find(code)
	if err != nil {
		if errors.Cause(err) == errors.New("redirect Not Found") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)
}

func (h *handler) getCodePathParameter(r *http.Request) (string, error) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 2 {
		return "", errors.New("Missing 'code' path parameter")
	}

	code := pathParts[1]
	_, err := strconv.Atoi(code)
	if err != nil {
		return "", errors.New("Missing 'code' path parameter")
	}

	fmt.Printf("GET request with code: %s\n", code)
	return code, nil
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	redirect, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.redirectService.Store(redirect)
	if err != nil {
		// TODO
		if errors.Cause(err) == errors.New("redirect Invalid") {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, err := h.serializer(contentType).Encode(redirect)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}
