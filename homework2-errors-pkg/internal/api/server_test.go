package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func decodeJSON[T any](t *testing.T, body *httptest.ResponseRecorder, dst *T) {
	t.Helper()
	require.NoError(t, json.NewDecoder(body.Body).Decode(dst))
}

func TestHandlers(t *testing.T) {
	h := New()

	t.Run("health", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)

		var resp map[string]string
		decodeJSON(t, rr, &resp)
		require.Equal(t, "ok", resp["status"])
	})

	t.Run("echo ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/echo?msg=ping", nil)
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)

		var resp map[string]string
		decodeJSON(t, rr, &resp)
		require.Equal(t, "ping", resp["echo"])
	})

	t.Run("echo missing param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/echo", nil)
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)

		var resp map[string]string
		decodeJSON(t, rr, &resp)
		require.Contains(t, resp["error"], "msg")
	})

	t.Run("item ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/item/1", nil)
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)

		var itm Item
		decodeJSON(t, rr, &itm)
		require.Equal(t, Item{ID: 1, Name: "foo"}, itm)
	})

	t.Run("item bad id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/item/abc", nil)
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("item not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/item/999", nil)
		rr := httptest.NewRecorder()

		h.ServeHTTP(rr, req)

		require.Equal(t, http.StatusNotFound, rr.Code)
	})
}
