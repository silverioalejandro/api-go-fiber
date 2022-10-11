package migration

import (
	"fmt"
	"go-fiber-gorm/database"
	"log"
	"github.com/silverioalejandro/api-go-fiber/model/entity"
)

func RunMigrations() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
