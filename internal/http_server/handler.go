package http_server

import (
	"encoding/json"
	"net/http"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", health)
	mux.HandleFunc("/", hello)
	return mux
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func hello(w http.ResponseWriter, _ *http.Request) {
	resp := struct{ Message string }{"Hello World"}
	_ = json.NewEncoder(w).Encode(resp)
}
