package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/get", get)
	fmt.Println("running in port 8081")
	http.ListenAndServe(":8081", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "get response from server")
	} else {
		fmt.Fprintf(w, fmt.Sprintf("%s response from server", r.Method))
	}
}
