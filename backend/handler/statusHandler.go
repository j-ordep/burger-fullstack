package handler

import (
	"backend-super-burger/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	service *service.StatusService
}

func NewStatusHandler(service *service.StatusService) *StatusHandler {
	return &StatusHandler{service: service}
}

func (h *StatusHandler) GetStatus(c echo.Context) error {
	status, err := h.service.GetStatus()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": "error ao buscar status",
		})
	}

	return c.JSON(http.StatusOK, status)
}