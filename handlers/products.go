package handlers

import (
	"go-microservices/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := &data.Product{}

	err := product.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "--Unable to unmarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

// ServeHTTP is OBSOLETE now since we use gorilla/mux to find out
//   which httpHandler we are going to continue on based on the request path,
//func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// ...
//}
