package input_fiber

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	enterprise_errors "github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		err := c.Query("error")
		switch err {
		case "notfound":
			return HandleError(c, &enterprise_errors.NotFoundError{Message: "resource not found"})
		case "conflict":
			return HandleError(c, &enterprise_errors.ConflictError{Message: "resource conflict"})
		case "default":
			return HandleError(c, errors.New("unexpected error"))
		default:
			return HandleError(c, nil)
		}
	})

	t.Run("should handle NotFoundError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test?error=notfound", nil)
		resp, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		expected := `{"message":"resource not found"}`
		assert.JSONEq(t, expected, string(body))
	})

	t.Run("should handle ConflictError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test?error=conflict", nil)
		resp, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusConflict, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		expected := `{"message":"resource conflict"}`
		assert.JSONEq(t, expected, string(body))
	})

	t.Run("should handle default error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test?error=default", nil)
		resp, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		expected := `{"message":"Internal server error"}`
		assert.JSONEq(t, expected, string(body))
	})

	t.Run("should handle nil error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		assert.Empty(t, bytes.TrimSpace(body))
	})
}
