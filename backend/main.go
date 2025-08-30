package main

import (
	"backend-super-burger/config"
	"backend-super-burger/db"
	"backend-super-burger/handler"
	"backend-super-burger/repository"
	"backend-super-burger/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.OpenConnection()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	ingredientesRepo := repository.NewIngredientesRepository(db)
	ingredientesService := service.NewIngredientesService(ingredientesRepo)
	ingredientesHandler := handler.NewIngredientesHandler(ingredientesService)

	statusRepo := repository.NewStatusRepository(db)
	statusService := service.NewStatusService(statusRepo)
	statusHandler := handler.NewStatusHandler(statusService)

	burgerRepo := repository.NewBurgerRepository(db)
	burgerService := service.NewBurgerService(burgerRepo)
	burgerHandler := handler.NewBurgerHandler(burgerService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/ingredientes", ingredientesHandler.GetIngredientes)

	e.GET("/status", statusHandler.GetStatus)

	e.GET("/burgers", burgerHandler.GetBurgers)
	e.POST("/burgers", burgerHandler.SaveBurger)
	e.PATCH("/burgers/:id", burgerHandler.UpdateStatusBurger)
	e.DELETE("/burgers/:id", burgerHandler.DeleteBurger)

	port := ":" + config.GetServerPort()
	e.Logger.Fatal(e.Start(port))
}