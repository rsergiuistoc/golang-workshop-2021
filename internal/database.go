package internal

import (
	"fmt"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabaseConn(c *Configuration) *gorm.DB {
	dns := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		c.DBHost, c.DBPort, c.DBUsername, c.DBName, c.DBPassword)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dns}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		panic(err)
	}

	return db
}