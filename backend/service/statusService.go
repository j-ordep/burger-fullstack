package service

import "backend-super-burger/domain"

type StatusService struct {
	repo domain.StatusRepository
}

func NewStatusService(repo domain.StatusRepository) *StatusService {
	return &StatusService{repo:repo}
}

func (s *StatusService) GetStatus() ([]*domain.Status, error){
	status, err := s.repo.GetStatus()
	if err != nil {
		return nil, err
	}

	return status, nil
}