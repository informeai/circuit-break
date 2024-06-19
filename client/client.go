package main

import (
	"fmt"
	"net/http"

	"github.com/informeai/circuit-break/config"
	"github.com/informeai/circuit-break/handler"
	"github.com/informeai/circuit-break/request"
)

func main() {
	circuitBreaker := config.CBConfig()
	request := request.NewHttpRequest(circuitBreaker, http.Client{}, "http://localhost:8081/api/v1/get")
	handler := handler.NewHandler(*request)

	http.HandleFunc("/api/v1/get", handler.Get)
	fmt.Println("running in port 8080")
	http.ListenAndServe(":8080", nil)
}
