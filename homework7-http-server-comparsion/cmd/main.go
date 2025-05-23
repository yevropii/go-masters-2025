package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/yevropii/go-masters-2025/homework7-http-server-comparsion/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("std http  -> :8080")
		if err := server.StartStd(":8080"); err != nil && ctx.Err() == nil {
			log.Fatalf("std server: %v", err)
		}
	}()

	go func() {
		log.Println("fasthttp -> :8081")
		if err := server.StartFast(":8081"); err != nil && ctx.Err() == nil {
			log.Fatalf("fasthttp server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down …")
	server.Stop()
	time.Sleep(200 * time.Millisecond)
}
