package db

import (
	"backend-super-burger/domain"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {

	err := db.AutoMigrate(
    &domain.Pao{},
    &domain.Carne{},
    &domain.Opcional{},
    &domain.Status{},
    &domain.Burger{},
    &domain.BurgerOpcional{},
)

	if err != nil {
		log.Printf("Erro ao executar migrações: %v", err)
		return err
	}

	return nil
}
