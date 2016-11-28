package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Price uint
}

func GetAll() int {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=gouser dbname=godb sslmode=disable password=gopass")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Product{})

	var product Product
	db.First(&product, 1)
	return int(product.Price)
}

func Create() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=gouser dbname=godb sslmode=disable password=gopass")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Price: 1000})
}
