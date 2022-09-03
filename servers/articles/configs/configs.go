package configs

import (
	helpers "go-grpc-micro/servers/articles/utils/generic"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Address  string
	Port     string
	Database struct {
		Name     string
		Port     string
		Username string
		Password string
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {

	err := godotenv.Load()
	if err != nil {
		helpers.Errorf("loading .env file server_articles" + err.Error())
		log.Print(&helpers.Buf)
	}

	var defaultConfig AppConfig
	defaultConfig.Address = os.Getenv("APP_ADDRESS")
	defaultConfig.Port = os.Getenv("ARTICLE_APP_PORT")
	defaultConfig.Database.Name = os.Getenv("POSTGRES_DB")
	defaultConfig.Database.Port = os.Getenv("POSTGRES_PORT")
	defaultConfig.Database.Username = os.Getenv("POSTGRES_USER")
	defaultConfig.Database.Password = os.Getenv("POSTGRES_PASSWORD")

	helpers.Server_user_port = os.Getenv("USER_APP_PORT")

	return &defaultConfig

}
