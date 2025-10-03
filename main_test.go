package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    w := httptest.NewRecorder()

    handler(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Fatalf("expected status 200, got %d", res.StatusCode)
    }

    body, _ := io.ReadAll(res.Body)
    expected := "Hello Jenkins CI/CD with Go + Docker!\n"
    if string(body) != expected {
        t.Fatalf("unexpected body: %q", string(body))
    }
}


