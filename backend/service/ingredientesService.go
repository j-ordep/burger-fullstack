package service

import (
	"backend-super-burger/domain"
)

type IngredientesService struct {
	repo domain.IngredientesRepository
}

func NewIngredientesService(repo domain.IngredientesRepository) *IngredientesService {
	return &IngredientesService{repo: repo}
}

func (s *IngredientesService) GetIngredientes() (*domain.IngredientesResponse, error) {

	paes, err := s.repo.GetPaes()
	if err != nil {
		return nil, err
	}

	carnes, err := s.repo.GetCarnes()
	if err != nil {
		return nil, err
	}

	opcionais, err := s.repo.GetOpcionais()
	if err != nil {
		return nil, err
	}

	return &domain.IngredientesResponse {
		Paes: paes,
		Carnes: carnes,
		Opcionais: opcionais,
	}, nil
	
}