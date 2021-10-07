package model

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv.Error())
	}
	
  database, err := gorm.Open(mysql.Open(os.Getenv("DB_URL")))

  if err != nil {
    panic(err.Error())
  }

  database.AutoMigrate(&Question{})
	database.AutoMigrate(&User{})

  DB = database
}