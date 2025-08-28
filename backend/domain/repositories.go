package domain

type IngredientesRepository interface {
	GetPaes() ([]*Pao, error)
	GetCarnes() ([]*Carne, error)
	GetOpcionais() ([]*Opcional, error)
}

type StatusRepository interface {
	GetStatus() ([]*Status, error)
}

type BurgerRepository interface {
	SaveBurger(burger *Burger) (int, error)
	GetStatusIdByName(statusName string) (int, error)
	GetBurgers() ([]*Burger, error)
	GetStatusTypeById(statusId int) (string, error)
	DeleteBurger(id int) error 
}