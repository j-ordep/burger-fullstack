package domain

type Status struct {
	Id int `json:"id" gorm:"primaryKey;column:id"`
    Tipo string `json:"tipo" gorm:"column:tipo"`
}

func (Status) TableName() string {
    return "status"
}