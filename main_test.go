package main

import (
	"net/http"
	"net/http/httptest"

	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

const (
	Status = "{\"Status\":\"OK\"}\n"
)
func TestRoot(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)

	if assert.NoError(t, rootHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, Hellow, rec.Body.String())
	}
}

func TestPing(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)

	if assert.NoError(t, pingHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, Status, rec.Body.String())
	}
}
