package domain

type Burger struct {
    Id        int      `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
    Nome      string   `json:"nome" gorm:"column:nome"`
    Pao       string   `json:"pao" gorm:"column:pao"`
    Carne     string   `json:"carne" gorm:"column:carne"`
    StatusId  int      `json:"status_id" gorm:"column:status_id"`
    Status    *Status  `json:"-" gorm:"foreignKey:StatusId"`
    Opcionais []string `json:"opcionais" gorm:"-"` // Tratado separadamente
}

func (Burger) TableName() string {
    return "pedidos"
}