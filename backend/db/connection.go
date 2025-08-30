package db

import (
	"backend-super-burger/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {
	conf := config.GetDB() 

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

	// Configurações de conexão
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)

    log.Println("Conexão com banco de dados estabelecida com sucesso")
    return db, nil
}
