package service

import (
	"backend-super-burger/domain"
)

type BurgerService struct {
	repo domain.BurgerRepository
}

func NewBurgerService(repo domain.BurgerRepository) *BurgerService {
	return &BurgerService{repo: repo}
}

func (s *BurgerService) SaveBurger(burger *domain.Burger) error {
	err := s.repo.SaveBurger(burger)
	if err != nil {
		return err
	}
	return nil
}