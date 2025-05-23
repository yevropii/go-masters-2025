package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFiboHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/fibo?N=10", nil)
	rec := httptest.NewRecorder()

	fiboHandler(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if got, want := strings.TrimSpace(string(body)), "55"; got != want {
		t.Fatalf("body = %q, want %q", got, want)
	}
}
