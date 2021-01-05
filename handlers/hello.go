package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello structu'ı httpHandler interface'ini implement edecek olan struct.
// HttpHandler oluşturmak için yapmamız gereken tek şey interface'i implement eden bir struc oluşturmak.
type Hello struct {
	l *log.Logger
}

// NewHello creates a Hello HttpHandler instance with given Logger object
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHttp is the method that satisfies HttpHandler Interface
// (h *Hello)											> Tip ismi
// ServeHTTP(rw http.ResponseWriter, r http.Request)	> Signature
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//log.Println("Hello World!")
	// Artık enjekte edilen logger objesini kullanıyoruz:
	h.l.Println("Hello World!")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops!", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Hello, %s", data)
}
