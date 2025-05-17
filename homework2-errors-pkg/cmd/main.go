package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yevropii/go-masters-2025/homework2-errors-pkg/internal/api"
)

func main() {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      api.New(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("HTTP сервер на %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка серевера: %v", err)
	}
}
