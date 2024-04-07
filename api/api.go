package api

import (
	"net/http"
	"os"
	"os/signal"
	"shop-test/model"
	"shop-test/pkg/log"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Api struct {
	e *echo.Echo
}

type ServiceApiIf interface {
	Run(server *http.Server, logger log.ILogger, debug bool) error
}

func NewApi(svc model.Controller) *Api {
	e := echo.New()

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	e.GET("/", apiHelthCheck(svc))


	e.POST("/item", ApiCreateItem(svc))
	e.GET("/item/:id", ApiGetItem(svc))
	e.GET("/item", ApiListItem(svc))

	return &Api{
		e: e,
	}
}

func (c *Api) Run(server *http.Server, logger log.ILogger) error {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		logger.Info("INF: Server shutdown form port ", server.Addr)
		os.Exit(0)
	}()

	c.e.HideBanner = true
	c.e.Debug = true
	c.e.Logger.SetOutput(NewNullWriter())
	return c.e.StartServer(server)
}
