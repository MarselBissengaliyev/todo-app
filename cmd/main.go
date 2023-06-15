package main

import (
	"log"

	"github.com/MarselBisengaliev/go-todo-app"
	"github.com/MarselBisengaliev/go-todo-app/cmd/pkg/handler"
	"github.com/MarselBisengaliev/go-todo-app/cmd/pkg/repository"
	"github.com/MarselBisengaliev/go-todo-app/cmd/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occur	ed while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}