package loadtest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRun(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	stats := Run(srv.URL, 4, 100)
	if stats.Failed != 0 {
		t.Fatalf("want 0 failed, got %d", stats.Failed)
	}
	if stats.Success != 100 {
		t.Fatalf("want 100 success, got %d", stats.Success)
	}
}
