package handler

import (
	"encoding/json"
	"net/http"

	"github.com/informeai/circuit-break/request"
)

func NewHandler(httpRequest request.HttpRequest) *Handler {
	return &Handler{httpRequest}
}

type Handler struct {
	httpRequest request.HttpRequest
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	res, err := h.httpRequest.GetToServerApp()
	data := map[string]string{
		"message": "message",
		"data":    string(res),
	}

	byteData, _ := json.Marshal(data)
	if err != nil {
		w.Write(byteData)
		return
	}
	w.Write(byteData)
}
