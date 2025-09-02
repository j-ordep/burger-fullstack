package domain

type Pao struct {
    Id   int    `json:"id" gorm:"primaryKey;column:id"`
    Tipo string `json:"tipo" gorm:"column:tipo"`
}

func (Pao) TableName() string {
    return "paes"
}

type Carne struct {
    Id   int    `json:"id" gorm:"primaryKey;column:id"`
    Tipo string `json:"tipo" gorm:"column:tipo"`
}

func (Carne) TableName() string {
    return "carnes"
}

type Opcional struct {
    Id   int    `json:"id" gorm:"primaryKey;column:id"`
    Tipo string `json:"tipo" gorm:"column:tipo"`
}

func (Opcional) TableName() string {
    return "opcionais"
}

// tabela de relacionamento
type BurgerOpcional struct {
    Id        int    `json:"id" gorm:"primaryKey;column:id"`
    PedidosId int    `json:"pedidos_id" gorm:"column:pedidos_id"`
    Opcional  string `json:"opcional" gorm:"column:opcional"`
}

func (BurgerOpcional) TableName() string {
    return "burger_opcionais"
}

type IngredientesResponse struct {
    Paes      []*Pao      `json:"paes"`
    Carnes    []*Carne    `json:"carnes"`
    Opcionais []*Opcional `json:"opcionais"`
}