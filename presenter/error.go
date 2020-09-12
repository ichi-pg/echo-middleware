package presenter

import (
	"github.com/labstack/echo/v4"
)

// ErrorPresenter はエラーレスポンスを抽象化します。
type ErrorPresenter interface {
	Response(echo.Context, error) error
}
