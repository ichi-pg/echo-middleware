package middleware

import (
	"fmt"
	"runtime"

	"github.com/ichi-pg/echo-middleware/contexts"
	"github.com/ichi-pg/echo-middleware/presenter"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Recover は例外をレスポンスに変換します。
func Recover(ep presenter.ErrorPresenter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					ep.Response(c, err)

					stack := make([]byte, 4096)
					length := runtime.Stack(stack, true)

					log := contexts.Logger(c)
					log.WithField("response", logrus.Fields{
						"status": c.Response().Status,
						"error":  fmt.Sprintf("%+v", err),
						"stack":  fmt.Sprintf("%s", stack[:length]),
					}).Error(accessLog)
				}
			}()
			return next(c)
		}
	}
}
