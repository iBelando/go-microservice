package commons

import (
	"log"

	"github.com/iBelando/go-microservice/models"
	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/test?charset=utf8")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando...")

	db.AutoMigrate(&models.Persona{})
}
