package api

import (
	// "io/ioutil"
	// "log"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
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
	code := chi.URLParam(r, "code")
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

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {}
