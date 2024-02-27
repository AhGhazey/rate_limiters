package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecovery_PanicRecovery(t *testing.T) {
	panicMessage := "Simulated panic"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(panicMessage)
	})

	rr := httptest.NewRecorder()

	recoveryMiddleware := Recovery(handler)
	recoveryMiddleware.ServeHTTP(rr, httptest.NewRequest("GET", "/test", nil))

	assert.Equal(t, http.StatusInternalServerError, rr.Code, "HTTP status code should be 500")
}

func TestRecovery_NextHandlerCalled(t *testing.T) {
	nextHandlerCalled := false

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextHandlerCalled = true
	})

	rr := httptest.NewRecorder()

	recoveryMiddleware := Recovery(handler)
	recoveryMiddleware.ServeHTTP(rr, httptest.NewRequest("GET", "/test", nil))

	// Use assert to simplify the assertion
	assert.True(t, nextHandlerCalled, "Expected the next http to be called, but it wasn't.")
	assert.Equal(t, http.StatusOK, rr.Code, "HTTP status code should be 200")
}
