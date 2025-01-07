package main

import (
	"os"
	project "school21_project1"
	"school21_project1/pkg/handler"
	"school21_project1/pkg/repository"
	"school21_project1/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error to initializing config %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error to loading godotenv: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBname:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed connect to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	server := new(project.Server)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("failed to run server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
