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
	UpdateStatusBurger(id int, statusId int) error
	DeleteBurger(id int) error 
	GetBurgers() ([]*Burger, error)
	GetBurgerById(id int) (*Burger, error)
}