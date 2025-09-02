package repository

import (
	"backend-super-burger/domain"

	"gorm.io/gorm"
)

type BurgerRepository struct {
	db *gorm.DB
}

func NewBurgerRepository(db *gorm.DB) *BurgerRepository {
	return &BurgerRepository{db: db}
}

func (r *BurgerRepository) SaveBurger(burger *domain.Burger) (int, error) {

	err := r.db.Create(burger).Error
	if err != nil {
		return 0, err
	}

	for _, opcional := range burger.Opcionais {
		burgerOpcional := domain.BurgerOpcional{
			PedidosId: burger.Id,
			Opcional:  opcional,
		}
		err := r.db.Create(&burgerOpcional).Error
		if err != nil {
			return 0, err
		}
	}

	return burger.Id, nil
}

func (r *BurgerRepository) GetBurgers() ([]*domain.Burger, error) {
	var burgers []*domain.Burger

	err := r.db.Preload("Status").Order("id").Find(&burgers).Error
	if err != nil {
		return nil, err
	}

	for _, burger := range burgers {
		var burgerOpcionais []domain.BurgerOpcional

		err := r.db.Where("pedidos_id = ?", burger.Id).Find(&burgerOpcionais).Error
		if err != nil {
			return nil, err
		}

		opcionais := []string{}
		for _, o := range burgerOpcionais {
			opcionais = append(opcionais, o.Opcional)
		}
		burger.Opcionais = opcionais
	}

	return burgers, nil
}

func (r *BurgerRepository) GetBurgerById(id int) (*domain.Burger, error) {
	var burger domain.Burger

	err := r.db.Preload("Status").First(&burger, id).Error
	if err != nil {
		return nil, err
	}

	var burgerOpcionais []domain.BurgerOpcional

	err = r.db.Where("pedidos_id = ?", id).Find(&burgerOpcionais).Error
	if err != nil {
		return nil, err
	}

	opcionais := []string{}
	for _, o := range burgerOpcionais {
		opcionais = append(opcionais, o.Opcional)
	}
	burger.Opcionais = opcionais

	return &burger, nil
}

func (r *BurgerRepository) UpdateStatusBurger(id int, statusId int) error {
	err := r.db.Model(&domain.Burger{}).Where("id = ?", id).Update("status_id", statusId).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BurgerRepository) DeleteBurger(id int) error {

	err := r.db.Delete(&domain.Burger{}, id).Error
	if err != nil {
		return err
	} 

	var count int64
	err = r.db.Model(&domain.Burger{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return r.db.Exec("ALTER SEQUENCE pedidos_id_seq RESTART WITH 1").Error
	}

	return nil
}