package domain

type Pao struct {
	Id int	
	Tipo string
}

type Carne struct {
	Id int	
	Tipo string
}

type Opcional struct {
	Id int	
	Tipo string
}

type IngredientesResponse struct {
    Paes      []*Pao      `json:"paes"`
    Carnes    []*Carne    `json:"carnes"`
    Opcionais []*Opcional `json:"opcionais"`
}