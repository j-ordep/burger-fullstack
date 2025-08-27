package domain

type Burger struct {
    Nome      string   `json:"nome"`
    Pao       string   `json:"pao"`
    Carne     string   `json:"carne"`
    Opcionais []string `json:"opcionais"`
    StatusId  int      `json:"status"`
}