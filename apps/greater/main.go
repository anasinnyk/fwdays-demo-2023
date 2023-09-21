package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
}

func main() {
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		var p Person

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		fmt.Fprintf(w, "Hello, %v", p.Name)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
