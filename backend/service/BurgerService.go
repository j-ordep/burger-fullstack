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

func (s *BurgerService) SaveBurger(burger *domain.Burger) (int, error) {
	burgerId, err := s.repo.SaveBurger(burger)
	if err != nil {
		return 0, err
	}
	return burgerId, nil
}

func (s *BurgerService) GetBurgers() ([]*domain.Burger, error) {
	burgers, err := s.repo.GetBurgers()
	if err != nil {
		return nil, err
	}

	return burgers, nil
}

func (s *BurgerService) UpdateStatusBurger(id int, statusId int) (*domain.Burger, error) {
	
	err := s.repo.UpdateStatusBurger(id, statusId)
	if err != nil {
		return nil, err
	}

	burger, err := s.repo.GetBurgerById(id)
	if err != nil {
		return nil, err
	}

	return burger, nil
}

func (s *BurgerService) DeleteBurgerById(id int) (error) {
	err := s.repo.DeleteBurger(id)
	if err != nil {
		return err
	}

	return nil
}