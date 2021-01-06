package main

import (
	"go-microservices/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	// önceki derste defaultServeMux'ı kullanıyorduk. Artık kendi oluşturduğumuz ServeMux'ı kullanıyoruz.
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/good-bye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
