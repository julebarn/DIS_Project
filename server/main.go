package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./build")))
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "hello world"}`))
		fmt.Println("API called")
	})

	http.ListenAndServe(":8080", nil)
}
