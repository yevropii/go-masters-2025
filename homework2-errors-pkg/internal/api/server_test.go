package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemHandler_NotFound(t *testing.T) {
	server := New()
	req := httptest.NewRequest(http.MethodGet, "/items/999", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("ожидался статус 404, получен %d", res.StatusCode)
	}

	var body struct {
		Code    int
		Message string
	}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("ошибка декодирования ответа: %v", err)
	}

	if body.Code != 1001 {
		t.Fatalf("ожидался код 1001, получен %d", body.Code)
	}

	if body.Message == "" {
		t.Fatal("ожидалось не пустое сообщение")
	}
}
