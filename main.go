package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/oznoztn/go-microservices/handlers"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/products", ph.GetProducts)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/product", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/api/product/update/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// The service won't block as i wrapped it up in a go function.
	// But that means it also going to immediately shut down
	go func() {
		// ListenAndServe bloklayacaktı. O yüzden go func içerisine aldık.
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// signal.Notify will broadcast a message on the channel whenever it receives Kill & Interrupt signals from the OS
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// Reading from a channel block until the messages available to be consumed
	issuedSignal := <-signalChannel
	l.Println("Received terminate, graceful shutdown.", issuedSignal)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
