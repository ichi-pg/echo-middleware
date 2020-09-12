package middleware

import (
	"net/http"
	"os"
	"strconv"

	"github.com/ichi-pg/echo-middleware/env"
	"github.com/labstack/echo/v4"
)

// DebugMode はデバッグ起動時のみアクセスを許可します。
func DebugMode() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			debugMode, err := strconv.ParseBool(os.Getenv(env.DebugMode))
			if err != nil {
				return err
			}
			if !debugMode {
				return echo.NewHTTPError(http.StatusForbidden, "403 Forbidden")
			}
			return next(c)
		}
	}
}
