package repository

import (
	"backend-super-burger/domain"
	"database/sql"
)

type BurgerRepository struct {
	db *sql.DB
}

func NewBurgerRepository(db *sql.DB) *BurgerRepository{
	return &BurgerRepository{db:db}
}

func (r *BurgerRepository) SaveBurger(burger *domain.Burger) (int, error) {
	var id int

	err := r.db.QueryRow(`
		INSERT INTO pedidos (nome, pao, carne, status_id) VALUES ($1, $2, $3, $4) RETURNING id`,
		burger.Nome, burger.Pao, burger.Carne, burger.StatusId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	for _, opcional := range burger.Opcionais {
		_, err = r.db.Exec(`
			INSERT INTO burger_opcionais (burger_id, opcional) VALUES ($1, $2)`,
			id, opcional,
		)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}