package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloWorld hello world と返すだけのメソッド
func HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
