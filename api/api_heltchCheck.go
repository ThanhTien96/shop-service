package api

import (
	"net/http"
	"shop-test/model"

	"github.com/labstack/echo/v4"
)

func apiHelthCheck(svc model.Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		svc.Logger().Info("checking app.")
		appHelth := model.AppHelth{
			Name:    "Shop Service",
			Version: svc.Config().App.Version,
			Status: "avalidable",
		}

		return c.JSON(http.StatusOK, JsonSuccess("check helth app.", appHelth))
	})
}
