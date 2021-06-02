package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandlerShouldBeOK(t *testing.T) {
	e := echo.New()
	e.GET("/health", Health)

	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}

func TestHealthHandlerShouldBeMethodNotAllowed(t *testing.T) {
	e := echo.New()
	e.GET("/health", Health)

	req := httptest.NewRequest("POST", "/health", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}
