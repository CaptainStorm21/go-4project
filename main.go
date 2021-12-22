package main

import (
	"log"
	"net/http"
)

func main() {
	// body
	http.HandleFunc("/",
					func(
					w http.ResponseWriter,
					r *http.Request) {

		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1{
			name = names[0]
		}
		w.Write([]byte("hello "+ name))
	})

	err := http.ListenAndServe(":3000", nil)

	if err !=nil {
		log.Fatal(err)
	}
}
