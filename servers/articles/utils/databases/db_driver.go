package utils

import (
	"fmt"
	configs "go-grpc-micro/servers/articles/configs"
	entities_article "go-grpc-micro/servers/articles/entities/articles"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		config.Address,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("error connect into dbs articles", err.Error())
	}
	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(entities_article.Article{})
}
