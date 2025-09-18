package main

import (
	"Lesson15/internal/config"
	"Lesson15/internal/controller"
	"Lesson15/internal/repository"
	"Lesson15/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

// @title OnlineShop API
// @contact.name OnlineShop API Service
// @contact.url http://test.com
// @contact.email test@test.com
func main() {
	cfg := config.LoadConfig()
	db, err := sqlx.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)

	if err = ctrl.RunServer(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
