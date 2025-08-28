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

	// o input (burgerDto) vem com status em fomato de string, logo passei essa func para buscar no db qual id do status (exemplo: id:1, tipo:"Solicitado")
	statusId, err := h.service.GetStatusIdByName(burgerDto.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "status_id não encontrado",
		})
	}

	burger := burgerDto.ToDomain(statusId)

	burgerId, err := h.service.SaveBurger(burger)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "impossivel adicionar pedido",
		})
	}

	return c.JSON(http.StatusOK, burgerId)
}

func (h *BurgerHandler) GetBurgers(c echo.Context) error {
	burgers, err := h.service.GetBurgers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "impossível listar pedidos",
		})
	}

	var outputDTOs []*dto.Output

	for _, burger := range burgers {
		statusType, err := h.service.GetStatusTypeById(burger.StatusId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "erro ao processar dados dos pedidos",
			})
		}

		outputDTO := dto.FromDomain(burger, statusType)
		outputDTOs = append(outputDTOs, outputDTO)
	}

	return c.JSON(http.StatusOK, outputDTOs)
}

func (h *BurgerHandler) UpdateStatusBurger(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	var body struct {
		Status string `json:"status"`
	}

	err = c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	newDomainBurger, err := h.service.UpdateStatusBurger(id, body.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string {
			"error": "não foi possível atualizar o pedido",
		})
	}

	newStatus, err := h.service.GetStatusTypeById(newDomainBurger.StatusId)

	output := dto.FromDomain(newDomainBurger, newStatus)

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