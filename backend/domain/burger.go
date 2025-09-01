package domain

type Burger struct {
    Id        int      `json:"id"`
    Nome      string   `json:"nome"`
    Pao       string   `json:"pao"`
    Carne     string   `json:"carne"`
    Opcionais []string `json:"opcionais"`
    StatusId  int      `json:"statusId"`
}