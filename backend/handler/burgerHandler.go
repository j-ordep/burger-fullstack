package handler

import (
	"backend-super-burger/dto"
	"backend-super-burger/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BurgerHandler struct {
	service *service.BurgerService
}

func NewBurgerHandler(service *service.BurgerService) *BurgerHandler {
	return &BurgerHandler{service: service}
}

func (h *BurgerHandler) SaveBurger(c echo.Context) error {
	var burgerDto dto.Input

	err := c.Bind(&burgerDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	burger := burgerDto.ToDomain()

	burgerId, err := h.service.SaveBurger(burger)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "impossivel adicionar pedido",
		})
	}

	return c.JSON(http.StatusOK, map[string]int {
		"id": burgerId,
	})
}

func (h *BurgerHandler) GetBurgers(c echo.Context) error {
	burgers, err := h.service.GetBurgers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "impossível listar pedidos",
		})
	}

	var outputs []*dto.Output

	for _, burger := range burgers {
		output := dto.FromDomain(burger)
		outputs = append(outputs, output)
	}

	return c.JSON(http.StatusOK, outputs)
}

func (h *BurgerHandler) UpdateStatusBurger(c echo.Context) error {
	idParam := c.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	var status struct {
		StatusId int `json:"status"`
	}
	
	err = c.Bind(&status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	newDomainBurger, err := h.service.UpdateStatusBurger(id, status.StatusId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string {
			"error": "não foi possível atualizar o pedido",
		})
	}

	output := dto.FromDomain(newDomainBurger)

	return c.JSON(http.StatusOK, output)
}

func (h *BurgerHandler) DeleteBurger(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	if err := h.service.DeleteBurgerById(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao deletar o pedido",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Pedido com id %d deletado com sucesso", id),
	})
}