package domain

type Pao struct {
	Id int `json:"id"`
	Tipo string `json:"tipo"`
}

type Carne struct {
	Id int `json:"id"`
	Tipo string `json:"tipo"`
}

type Opcional struct {
	Id int `json:"id"`
	Tipo string `json:"tipo"`
}

type IngredientesResponse struct {
    Paes      []*Pao      `json:"paes"`
    Carnes    []*Carne    `json:"carnes"`
    Opcionais []*Opcional `json:"opcionais"`
}