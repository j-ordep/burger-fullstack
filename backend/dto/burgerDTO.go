package dto

import "backend-super-burger/domain"

type Input struct {
	Nome      string   `json:"nome"`
	Pao       string   `json:"pao"`
	Carne     string   `json:"carne"`
	Opcionais []string `json:"opcionais"`
	StatusId  int      `json:"statusId"`
}

type Output struct {
    Id        int      `json:"id"`
    Nome      string   `json:"nome"`
    Pao       string   `json:"pao"`
    Carne     string   `json:"carne"`
    Opcionais []string `json:"opcionais"`
	StatusId  int      `json:"statusId"`
}

func (dto *Input) ToDomain() *domain.Burger {
	return &domain.Burger{
		Nome:      dto.Nome,
		Pao:       dto.Pao,
		Carne:     dto.Carne,
		Opcionais: dto.Opcionais,
		StatusId:  dto.StatusId,
	}
}

func FromDomain(burger *domain.Burger) *Output {
	return &Output{
		Id:        burger.Id,
    	Nome:      burger.Nome,
        Pao:       burger.Pao,
        Carne:     burger.Carne,
        Opcionais: burger.Opcionais,
        StatusId:  burger.StatusId,
    }
}	