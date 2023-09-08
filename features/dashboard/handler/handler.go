package handler

import (
	"immersive_project/klp3/features/dashboard"
	"immersive_project/klp3/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	DashboardService dashboard.DashboardServiceInterface
}

func New(service dashboard.DashboardServiceInterface) *DashboardHandler {
	return &DashboardHandler{
		DashboardService: service,
	}
}

func (handler *DashboardHandler) GetData(c echo.Context) error {
	result, err := handler.DashboardService.GetData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, resource not valid", nil))
	}
	var dataResp = DataToResponse(result)
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", dataResp))
}
