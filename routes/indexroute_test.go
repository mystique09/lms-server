package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/", strings.NewReader(userRequestJson))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testRoute.indexRoute(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Welcome! This is my backend API for my Class Management System personal project.", rec.Body.String())
	}
}
