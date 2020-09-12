package middleware

import (
	"net/http"

	"github.com/ichi-pg/echo-middleware/repository"
	"github.com/labstack/echo/v4"
)

// Maintenance はサーバーメンテナンス中の場合にアクセスを拒否します。
func Maintenance(repo repository.MaintenanceRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if repo.Active() {
				return echo.NewHTTPError(http.StatusServiceUnavailable, repo.Message())
			}
			return next(c)
		}
	}
}
