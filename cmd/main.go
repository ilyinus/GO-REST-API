package main

import (
	"context"
	"github.com/ilyinus/go-rest-api"
	"github.com/ilyinus/go-rest-api/pkg/handlers"
	"github.com/ilyinus/go-rest-api/pkg/repositories"
	"github.com/ilyinus/go-rest-api/pkg/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Go-rest-API
// @version 1.0
// @description API Server

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repositories.NewPostgresDB(repositories.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.SSLmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	httpHandlers := handlers.NewHandler(service)

	srv := new(rest.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), httpHandlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while runnig http server: %s", err.Error())
		}
	}()

	logrus.Print("App has been started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
