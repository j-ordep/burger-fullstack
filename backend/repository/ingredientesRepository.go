package repository

import (
	"backend-super-burger/domain"

	"gorm.io/gorm"
)

type IngredientesRepository struct {
	db *gorm.DB
}

func NewIngredientesRepository(db *gorm.DB) *IngredientesRepository {
	return &IngredientesRepository{db: db}
}

func (r *IngredientesRepository) GetPaes() ([]*domain.Pao, error) {
	var paes []*domain.Pao

	err := r.db.Find(&paes).Error
	if err != nil {
		return nil, err
	}

	return paes, nil
}

func (r *IngredientesRepository) GetCarnes() ([]*domain.Carne, error) {
	var carnes []*domain.Carne

	err := r.db.Find(&carnes).Error
	if err != nil {
		return nil, err
	}

	return carnes, nil
}

func (r *IngredientesRepository) GetOpcionais() ([]*domain.Opcional, error) {
	var opcionais []*domain.Opcional

	err := r.db.Find(&opcionais).Error
	if err != nil {
		return nil, err
	}

	return opcionais, nil
}