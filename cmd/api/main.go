package main

import (
	"fmt"

	"targon-oracle/internal/setup"
	"targon-oracle/internal/shared"

	"github.com/aidarkhanov/nanoid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed init logger")
	}
	log := logger.Sugar()
	core, errs := setup.CreateCore()
	if errs != nil {
		panic(fmt.Sprintf("Failed creating core: %s", errs))
	}
	defer core.Shutdown()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId, _ := nanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 28)
			logger := log.With(
				"request_id", "req_"+reqId,
			)

			cc := &shared.Context{Context: c, Log: logger, Reqid: reqId, Core: core}
			return next(cc)
		}
	})
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			defer func() {
				_ = log.Sync()
			}()
			log.Errorw("Api Panic", "error", err.Error())
			return c.String(500, "Internal Server Error")
		},
	}))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	e.GET("/prices/h200/average", func(c echo.Context) error {
		// Placeholder response
		return c.JSON(200, map[string]interface{}{
			"gpu_type":      "h200",
			"average_price": 0.0, // Placeholder value
		})
	})

	e.Logger.Fatal(e.Start(":443"))
}
