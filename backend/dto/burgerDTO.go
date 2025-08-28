package dto

import "backend-super-burger/domain"

type Input struct {
	Nome      string   `json:"nome"`
	Pao       string   `json:"pao"`
	Carne     string   `json:"carne"`
	Opcionais []string `json:"opcionais"`
	Status    string   `json:"status"`
}

type Output struct {
    Id        int      `json:"id"`
    Nome      string   `json:"nome"`
    Pao       string   `json:"pao"`
    Carne     string   `json:"carne"`
    Opcionais []string `json:"opcionais"`
	Status    string   `json:"status"`
}

func (dto *Input) ToDomain(statusId int) *domain.Burger {
	return &domain.Burger{
		Nome:      dto.Nome,
		Pao:       dto.Pao,
		Carne:     dto.Carne,
		Opcionais: dto.Opcionais,
		StatusId:  statusId,
	}
}

func FromDomain(burger *domain.Burger, status string) *Output {
	return &Output{
		Id:        burger.Id,
    	Nome:      burger.Nome,
        Pao:       burger.Pao,
        Carne:     burger.Carne,
        Opcionais: burger.Opcionais,
        Status:    status,
    }
}	