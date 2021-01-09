package handlers

import (
	"context"
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

	//prod := &data.Product{}
	//err = prod.FromJson(r.Body)
	// Since we use middleware, we can fetch the product from the request context
	// Value returns an interface, so we need to cast this to Product
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
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

	//prod := &data.Product{}
	//err = prod.FromJson(r.Body)
	// Since we use middleware, we can fetch the product from the request context
	// Value returns an interface, so we need to cast this to Product
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal JSON", http.StatusInternalServerError)
			return
		}

		// add the product to the context of the next middleware
		// Key string tipinde.. Fakat direkt string vermek yerine struct vermek daha iyi ve genel bir yaklaşım.
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		p.l.Println("The product is valid. Calling the next middleware...")
		next.ServeHTTP(rw, r)
	})
}

// ServeHTTP is OBSOLETE now since we use gorilla/mux to find out
//   which httpHandler we are going to continue on based on the request path,
//func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// ...
//}
