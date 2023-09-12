package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/jp-chl/test-go-clean-architecture/api"
	dr "github.com/jp-chl/test-go-clean-architecture/domain/repository"
	mr "github.com/jp-chl/test-go-clean-architecture/repository"
	uc "github.com/jp-chl/test-go-clean-architecture/usecase"
)

func main() {
	repo := getDB()
	service := uc.NewRedirectService(repo)
	handler := api.NewHandler(service)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.Get(w, r)
		case http.MethodPost:
			handler.Post(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := httpPort()
	server := &http.Server{
		Addr: ":" + port,
	}

	_errors := make(chan error, 2)

	go func() {
		err := server.ListenAndServe()
		fmt.Println("Listening on port ", port)
		if err != nil {
			_errors <- err
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		_errors <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s\n", <-_errors)
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return port
}

func getDB() dr.RedirectRepository {
	switch os.Getenv("DB") {
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongoDB := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
