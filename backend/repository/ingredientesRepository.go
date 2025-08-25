package repository

import (
	"backend-super-burger/domain"
	"database/sql"
)

type IngredientesRepository struct {
	db *sql.DB
}

func NewIngredientesRepository(db *sql.DB) *IngredientesRepository {
	return &IngredientesRepository{db:db}
}


func (r *IngredientesRepository) GetPaes() ([]*domain.Pao, error) {
	var paes []*domain.Pao

	rows, err := r.db.Query("SELECT id, tipo FROM paes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &domain.Pao{}
		err := rows.Scan(&p.Id, &p.Tipo)
		if err != nil {
			return nil, err
		}
		paes = append(paes, p)
	}

	return paes, nil
}

func (r *IngredientesRepository) GetCarnes() ([]*domain.Carne, error) {
	var carnes []*domain.Carne

	rows, err := r.db.Query("SELECT id, tipo FROM carnes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &domain.Carne{}
		err := rows.Scan(&c.Id, &c.Tipo)
		if err != nil {
			return nil, err
		} 
		carnes = append(carnes, c)
	}

	return carnes, nil
}

func (r *IngredientesRepository) GetOpcionais() ([]*domain.Opcional, error) {
	var opcionais []*domain.Opcional

	rows, err := r.db.Query("SELECT id, tipo FROM opcionais")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		o := &domain.Opcional{}
		err := rows.Scan(&o.Id, &o.Tipo)
		if err != nil {
			return nil, err
		}
		opcionais = append(opcionais, o)
	}

	return opcionais, nil
}

