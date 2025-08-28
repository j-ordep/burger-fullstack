package repository

import (
	"backend-super-burger/domain"
	"database/sql"
)

type BurgerRepository struct {
	db *sql.DB
}

func NewBurgerRepository(db *sql.DB) *BurgerRepository {
	return &BurgerRepository{db: db}
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
			INSERT INTO burger_opcionais (pedidos_id, opcional) VALUES ($1, $2)`,
			id, opcional,
		)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *BurgerRepository) GetBurgers() ([]*domain.Burger, error) {
	rows, err := r.db.Query(`
		SELECT p.id, p.nome, p.pao, p.carne, p.status_id, s.tipo
		FROM pedidos p
		INNER JOIN status s ON p.status_id = s.id
		ORDER BY p.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	burgersMap := make(map[int]*domain.Burger)
	statusMap := make(map[int]string)

	for rows.Next() {
		var burger domain.Burger
		var statusTipo string

		err := rows.Scan(&burger.Id, &burger.Nome, &burger.Pao, &burger.Carne, &burger.StatusId, &statusTipo)
		if err != nil {
			return nil, err
		}

		burger.Opcionais = []string{} // Inicializa slice vazio para opcionais
		burgersMap[burger.Id] = &burger
		statusMap[burger.StatusId] = statusTipo
	}

	opcionaisRows, err := r.db.Query(`
		SELECT pedidos_id, opcional
		FROM burger_opcionais
		ORDER BY pedidos_id
	`)
	if err != nil {
		return nil, err
	}
	defer opcionaisRows.Close()

	for opcionaisRows.Next() {
		var pedidoId int
		var opcional string

		err := opcionaisRows.Scan(&pedidoId, &opcional)
		if err != nil {
			return nil, err
		}

		if burger, ok := burgersMap[pedidoId]; ok {
			burger.Opcionais = append(burger.Opcionais, opcional)
		}
	}

	// Converte o mapa para um slice para retornar
	var burgers []*domain.Burger

	for _, burger := range burgersMap {
		burgers = append(burgers, burger)
	}

	return burgers, nil
}

func (r *BurgerRepository) GetBurgerById(id int) (*domain.Burger, error) {

	row := r.db.QueryRow(`SELECT id, nome, pao, carne, status_id FROM pedidos WHERE id = $1`, id)

	var burger domain.Burger
	err := row.Scan(&burger.Id, &burger.Nome, &burger.Pao, &burger.Carne, &burger.StatusId)
	if err != nil {
		return nil, err
	}

	burger.Opcionais = []string{}

	opcRows, err := r.db.Query("SELECT opcional FROM burger_opcionais WHERE pedidos_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer opcRows.Close()

	for opcRows.Next() {
		var opcional string
		if err := opcRows.Scan(&opcional); err != nil {
			return nil, err
		}
		burger.Opcionais = append(burger.Opcionais, opcional)
	}

	return &burger, nil
}

func (r *BurgerRepository) UpdateStatusBurger(id int, statusId int) error {
	_, err := r.db.Exec("UPDATE pedidos SET status_id = $1 WHERE id = $2", statusId, id)
	return err
}

func (r *BurgerRepository) DeleteBurger(id int) error {

	_, err := r.db.Exec(`DELETE FROM pedidos WHERE id = $1`, id)
	if err != nil {
		return err
	}

	// Agora checa se a tabela ficou vazia
	var count int
	err = r.db.QueryRow(`SELECT COUNT(*) FROM pedidos`).Scan(&count)
	if err != nil {
		return err
	}

	// Se não tem mais pedidos, reseta a sequence
	if count == 0 {
		_, err = r.db.Exec(`ALTER SEQUENCE pedidos_id_seq RESTART WITH 1`)
		if err != nil {
			return err
		}
	}

	return nil
}

// funções auxiliares

func (r *BurgerRepository) GetStatusIdByName(statusName string) (int, error) {
	var statusId int
	err := r.db.QueryRow(`SELECT id FROM status WHERE tipo = $1`, statusName).Scan(&statusId)
	return statusId, err
}

func (r *BurgerRepository) GetStatusTypeById(statusId int) (string, error) {
	var statusType string
	err := r.db.QueryRow(`SELECT tipo FROM status WHERE id = $1`, statusId).Scan(&statusType)
	return statusType, err
}
