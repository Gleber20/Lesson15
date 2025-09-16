package main

import (
	"Lesson15/internal/controller"
	"Lesson15/internal/repository"
	"Lesson15/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=Simuve39 dbname=onlineshop sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)

	if err = ctrl.RunServer(":8080"); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
