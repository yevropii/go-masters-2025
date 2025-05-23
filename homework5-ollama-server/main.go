package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	defaultAddr = ":8080"
	baseURL     = "http://localhost:11434"
	model       = "llama3"
)

type generateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type generateResponse struct {
	Response string `json:"response"`
}

func handler(client *http.Client, baseURL, model string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		if q == "" {
			http.Error(w, "`q` is empty", http.StatusBadRequest)
			return
		}

		reqBody, _ := json.Marshal(generateRequest{Model: model, Prompt: q, Stream: false})
		ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/api/generate", bytes.NewReader(reqBody))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			io.Copy(w, resp.Body)
			w.WriteHeader(resp.StatusCode)
			return
		}

		var gr generateResponse
		if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte(gr.Response))
	}
}

func main() {
	mux := http.NewServeMux()
	client := &http.Client{Timeout: 0}
	mux.HandleFunc("/", handler(client, baseURL, model))

	log.Printf("listening on %s, forwarding to %s (%s)", defaultAddr, baseURL, model)
	log.Fatal(http.ListenAndServe(defaultAddr, mux))
}
