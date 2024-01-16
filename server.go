package server

import "net/http"

func NewServer() *http.Server {
	return &http.Server{
		Addr: "0.0.0.0:8080",
	}
}
