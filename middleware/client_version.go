package middleware

import (
	"net/http"
	"strconv"

	"github.com/ichi-pg/echo-middleware/header"
	"github.com/ichi-pg/echo-middleware/repository"
	"github.com/labstack/echo/v4"
)

// ClientVersion は強制アップデートバージョン未満のアクセスを拒否します。
func ClientVersion(repo repository.ClientVersionRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			h := c.Request().Header
			pf := h.Get(header.Platform)
			version, err := strconv.Atoi(h.Get(header.ClientVersion))
			if err != nil {
				return err
			}
			if version < repo.Version(pf) {
				return echo.NewHTTPError(http.StatusUpgradeRequired, repo.Version(pf))
			}
			return next(c)
		}
	}
}
