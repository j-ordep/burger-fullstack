package repository

import (
	"backend-super-burger/domain"

	"gorm.io/gorm"
)

type StatusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *StatusRepository {
	return &StatusRepository{db: db}
}

func (r *StatusRepository) GetStatus() ([]*domain.Status, error) {
	var status []*domain.Status

	err := r.db.Find(&status).Error
	if err != nil {
		return nil, err
	}

	return status, nil
}