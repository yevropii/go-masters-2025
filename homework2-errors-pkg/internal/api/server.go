package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/yevropii/go-masters-2025/homework2-errors-pkg/internal/errs"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var store = map[int]Item{
	1: {ID: 1, Name: "foo"},
	2: {ID: 2, Name: "bar"},
}

func New() http.Handler {
	r := chi.NewRouter()
	r.Get("/echo", echoHandler)
	r.Get("/health", healthHandler)
	r.Get("/item/{id}", itemHandler)

	return r
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		respondError(w, errs.NewBadRequest("отсутствует параметр msg"))
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"echo": msg})
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	idSrt := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSrt)
	if err != nil || id <= 0 {
		respondError(w, errs.NewBadRequest("некорректный идентификатор id %q", idSrt))
		return
	}

	item, ok := store[id]
	if !ok {
		respondError(w, errs.NewNotFound("элемент с ID: %q не найден", id))
		return
	}

	writeJSON(w, http.StatusOK, item)
}

func respondError(w http.ResponseWriter, err error) {
	if httpErr, ok := err.(errs.HTTPError); ok {
		writeJSON(w, httpErr.StatusCode(), map[string]string{"error": httpErr.Error()})
		return
	}

	writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
