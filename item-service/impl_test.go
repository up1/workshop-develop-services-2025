package api_test

import (
	"api"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetItems(t *testing.T) {
	// Setup
	e := echo.New()
	server := api.NewServer()
	req := httptest.NewRequest(http.MethodGet, "/api/items", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, server.GetItems(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var items []api.Item
		err := json.Unmarshal(rec.Body.Bytes(), &items)
		assert.NoError(t, err)
		assert.Len(t, items, 2)
		assert.Equal(t, 1, items[0].Id)
		assert.Equal(t, "Item 1", items[0].Name)
		assert.Equal(t, 2, items[1].Id)
		assert.Equal(t, "Item 2", items[1].Name)
	}
}
