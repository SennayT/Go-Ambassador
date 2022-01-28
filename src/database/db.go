package database

import (
	"ambassador/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var sqlErr error

	DB, sqlErr = gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	if sqlErr != nil {
		panic(sqlErr.Error())
	}

}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
	DB.AutoMigrate(models.Product{})
}
