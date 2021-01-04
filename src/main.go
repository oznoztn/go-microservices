package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 127.0.0.1:9090 diyebilirdim.
	// Ben lokal makinemdeki bütün IP'lerin 9090 portuna bind ediyorum

	// HttpHandler'lar ile gelen Http Request objelerini işleyebilirim.

	// Tanımladığım httpHandler, DefaultServeMux üzerindeki "/" adresine (path) map'leniyor.
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// *http.Request ile de-reference etmeye gerek yok.
		// Onun yerine method signature'da "*http.Request" yerine "r*http.Request" demen yeterli.
		// request değişkenine de-reference ediyoruz.
		log.Println("Hello World!")

		// request.Body bir Stream.
		// ioutil.ReadAll ile bunu okuyabilirim.
		// bu metot iki şey dönebilir data ve varsa error, dolayısıyla:
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops!"))
			// return

			// Üsttek işlemler için kısayol metodu mevcut:
			http.Error(rw, "Oops!", http.StatusBadRequest)
		}

		// 		TEST: $ curl -v -d 'i'm ozan' localhost:9090
		log.Printf("Request Body: %s\n", data)

		fmt.Fprintf(rw, "Hello, %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good bye World!")
	})

	// nil verdiğimden dolayı defaultServeMux kullanılıyor.
	http.ListenAndServe(":9090", nil)

	// çalıştığını yeni bir bash terminalinde curl ile şöyle test edebilirsin
	// 		curl -v localhost:9090

}

// go build -o main.exe && ./main.exe
