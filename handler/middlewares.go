package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ContentTypeJson checks that the requests have the Content-Type header set to "application/json".
// This helps against CSRF attacks.
func ContentTypeJson(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")
		if contentType != "application/json" {
			return c.JSON(http.StatusBadRequest, jsonHTTPResponse{false, "只允许 JSON 请求"})
		}

		return next(c)
	}
}
