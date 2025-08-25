package repository

import (
	"backend-super-burger/domain"
	"database/sql"
)

type StatusRepository struct {
	db *sql.DB
}

func NewStatusRepository(db *sql.DB) *StatusRepository {
	return &StatusRepository{db: db}
}

func (r *StatusRepository) GetStatus() ([]*domain.Status, error) {
	var status []*domain.Status

	rows, err := r.db.Query("SELECT id, tipo FROM status")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := &domain.Status{}
		err := rows.Scan(&s.Id, &s.Tipo)
		if err != nil {
			return nil, err
		}
		status = append(status, s)
	}

	return status, nil
}