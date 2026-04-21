package util

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func NewResponse(status int, message string, data any) *Response {
	return &Response {
		status,
		message,
		data,
	}
}

func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	log.Printf("%d, %s", r.Status, r.Message)
	
	var err = json.NewEncoder(w).Encode(r)
	if err != nil {
		fmt.Printf("error encoding response struct\n")
	}
}