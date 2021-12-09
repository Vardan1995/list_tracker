package database

import (
	"fmt"

	"github.com/Vardan1995/list_tracker/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&entity.User{})
	DB.AutoMigrate(&entity.Filter{})
	DB.AutoMigrate(&entity.Archive{})
	fmt.Println("Database Migrated")
}
