package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_golang_api?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	//db.AutoMigrate(&entities.Division{})
	//db.AutoMigrate(&entities.Group{})
	//db.AutoMigrate(&entities.Person{})
	//db.AutoMigrate(&entities.PlaceOfBirth{})
	//db.AutoMigrate(&entities.Address{})

	DB = db
	log.Println("Database connected...")
}
