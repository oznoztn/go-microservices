package main

import (
	"go-microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	// önceki derste defaultServeMux'ı kullanıyorduk. Artık kendi oluşturduğumuz ServeMux'ı kullanıyoruz.
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/good-bye", gh)

	http.ListenAndServe(":9090", sm)
}
