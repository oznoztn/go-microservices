package handlers

import (
	"go-microservices/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := &data.Product{}

	err := product.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "--Unable to unmarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	// I cannot define an id field in the method signature to just get that
	// I need to get the product Id from the request context
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to parse Id field", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Product", id)
	prod := &data.Product{}

	err = prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusInternalServerError)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}

// ServeHTTP is OBSOLETE now since we use gorilla/mux to find out
//   which httpHandler we are going to continue on based on the request path,
//func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// ...
//}
