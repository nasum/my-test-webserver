package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// HelloWorld hello world と返すだけのメソッド
func HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		output := fmt.Sprintf("%#v", c.Request().Header)

		os.Stdout.Write([]byte(output + "\n"))
		return c.String(http.StatusOK, "Hello World")
	}
}
