package handler

import (
	"immersive_project/klp3/features/statuses"
	"immersive_project/klp3/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	StatusService statuses.StatusServiceInterface
}

func New(service statuses.StatusServiceInterface) *StatusHandler {
	return &StatusHandler{
		StatusService: service,
	}
}

func (handler *StatusHandler) GetAll(c echo.Context) error {
	result, err := handler.StatusService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	var statusResp []StatusResponse
	for _, v := range result {
		var status = StatusEntityToResponse(v)
		statusResp = append(statusResp, status)
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", statusResp))
}
