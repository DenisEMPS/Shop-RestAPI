package main

import (
	"context"
	"os"
	"os/signal"
	project "school21_project1"
	"school21_project1/pkg/handler"
	"school21_project1/pkg/repository"
	"school21_project1/pkg/service"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Online-Shop API
// @version 1.0
// @description API server for online-shop Application

// @host localhost:8000
// @BasePath /api/v1

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

	go func() {
		if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("failed to run server: %s", err.Error())
		}
	}()

	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Printf("error ocured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Printf("error ocured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
