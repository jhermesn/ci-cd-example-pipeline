package handler_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jhermesn/ci-cd-example-pipeline/handler"
)

func newCountRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/count", handler.Count)
	return r
}

func TestCount_HappyPath(t *testing.T) {
	r := newCountRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/count?text=hello+world", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if w.Body.String() != `{"count":2}` {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestCount_MultipleSpaces(t *testing.T) {
	r := newCountRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/count", nil)
	q := url.Values{}
	q.Set("text", "  foo   bar  baz  ")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if w.Body.String() != `{"count":3}` {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestCount_EmptyText(t *testing.T) {
	r := newCountRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/count?text=", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestCount_MissingParam(t *testing.T) {
	r := newCountRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/count", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
