package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, world!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request!", http.StatusBadRequest)
			return
		}
		log.Println(string(d))
		fmt.Fprintf(w, "Hello, %s!", string(d))
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye, world!")
		w.Write([]byte("Goodbye, world!"))
	})

	http.ListenAndServe(":9090", nil)
}
