package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":20100", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	l := r.ContentLength
	body := make([]byte, l)

	_, _ = r.Body.Read(body)
	_, _ = w.Write(body)
	fmt.Println(r.Method, r.URL, string(body))
}
