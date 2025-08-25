package handler

import (
	"backend-super-burger/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IngredientesHandler struct {
	service *service.IngredientesService
}

func NewIngredientesHandler(service *service.IngredientesService) *IngredientesHandler {
	return &IngredientesHandler{service: service}
}

func (h *IngredientesHandler) GetIngredientes(c echo.Context) error {
	ingredientes, err := h.service.GetIngredientes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": "Erro ao buscar Ingredientes",
		})
	}

	return c.JSON(http.StatusOK, ingredientes)
}