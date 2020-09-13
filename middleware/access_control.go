package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/ichi-pg/golang-middleware/env"
)

// AccessControl はIPホワイトリスト以外のアクセスを拒否します。
func AccessControl() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !strings.Contains(os.Getenv(env.WhiteIP), c.RealIP()) {
				return echo.NewHTTPError(http.StatusForbidden, "403 Forbidden")
			}
			return next(c)
		}
	}
}
