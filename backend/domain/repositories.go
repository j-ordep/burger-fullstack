package domain

type IngredientesRepository interface {
	GetPaes() ([]*Pao, error)
	GetCarnes() ([]*Carne, error)
	GetOpcionais() ([]*Opcional, error)
}

type StatusRepository interface {
	GetStatus() ([]*Status, error)
}