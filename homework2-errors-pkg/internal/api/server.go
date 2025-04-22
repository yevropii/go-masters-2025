package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/yevropii/go-masters-2025/homework2-errors-pkg/internal/errs"
)

// Сущность для наглядности
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Хранилище
var store map[int]Item

// Маршрутизатор, возвращает http.Handler
func New() http.Handler {
	r := chi.NewRouter()

	r.Get("/ok", healthHandler)
	r.Get("/echo", echoHandler)
	r.Get("/items/{id}", itemHandler)

	return r
}

// Handler для проверки статуса сервера
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Handler для возврата json с отправленным msg
func echoHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		respondErr(w, errs.New("BAD_REQUEST"))
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"echo": msg})
}

// Handler для запроса json с данными item по id
func itemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		respondErr(w, errs.New("BAD_REQUEST"))
		return
	}

	item, ok := store[id]
	if !ok {
		respondErr(w, errs.New("NOT_FOUND"))
		return
	}
	writeJSON(w, http.StatusOK, item)
}

// Mapper из AppError в Response
func respondErr(w http.ResponseWriter, appErr errs.AppError) {
	resp := struct {
		Code    int
		Message string
	}{
		Code:    appErr.Code(),
		Message: appErr.Message(),
	}

	writeJSON(w, appErr.StatusCode(), resp)
}

// Cериализует структуру в JSON
func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
