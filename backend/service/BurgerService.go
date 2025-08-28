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

func (s *BurgerService) DeleteBurgerById(id int) (error) {
	err := s.repo.DeleteBurger(id)
	if err != nil {
		return err
	}

	return nil
}

// funções auxiliares


func (s *BurgerService) GetStatusIdByName(status string) (int, error) {
	statusId, err := s.repo.GetStatusIdByName(status)
	if err != nil {
		return 0, err
	}

	return statusId, nil
}

func (s *BurgerService) GetStatusTypeById(statusId int) (string, error) {
	statustype, err := s.repo.GetStatusTypeById(statusId)
	if err != nil {
		return "", err
	}

	return statustype, nil
}