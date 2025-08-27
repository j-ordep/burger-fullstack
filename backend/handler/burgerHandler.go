package handler

import (
	"backend-super-burger/domain"
	"backend-super-burger/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BurgerHandler struct {
	service *service.BurgerService
}

func NewBurgerHandler(service *service.BurgerService) *BurgerHandler {
	return &BurgerHandler{service: service}
}

func (h *BurgerHandler) SaveBurger(c echo.Context) error {
	var burger *domain.Burger 
	err := c.Bind(&burger)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string {
			"error" : "JSON inválido",
		})
	} 

	err = h.service.SaveBurger(burger)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)

}