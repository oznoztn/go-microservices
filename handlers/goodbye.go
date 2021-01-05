package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GoodBye return GoodBye struct.
type GoodBye struct {
	l *log.Logger
}

// NewGoodBye creates a GoodBye HttpHandler instance with given Logger object
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

// ServeHttp is cool
func (gb *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	gb.l.Println("GoodBye HttpHandler was triggered.")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops!", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Good bye my dear friend %s", d)
}
