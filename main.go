package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/jp-chl/test-go-clean-architecture/api"
	dr "github.com/jp-chl/test-go-clean-architecture/domain/repository"
	mr "github.com/jp-chl/test-go-clean-architecture/repository"
	uc "github.com/jp-chl/test-go-clean-architecture/usecase"
)

func main() {
	repo := getDB()
	service := uc.NewRedirectService(repo)
	handler := api.NewHandler(service)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	fmt.Println("======= XXX ========")

	router.Get("/{code}", handler.Get)
	router.Post("/", handler.Post)
	fmt.Println("======= XXX ========")

	_errors := make(chan error, 2)
	fmt.Println("======= XXX ========")
	go func() {
		port := httpPort()
		fmt.Println("======= XXX ========")
		fmt.Println("Listenening on port ", port)
		http.ListenAndServe(port, router)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		fmt.Println("======= XXX ========")
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
	return fmt.Sprintf(":%s", port)
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
