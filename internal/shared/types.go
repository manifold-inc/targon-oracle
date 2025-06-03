package shared

import (
	"targon-oracle/internal/setup"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Context struct {
	echo.Context
	Log   *zap.SugaredLogger
	Reqid string
	Core  *setup.Core
}
