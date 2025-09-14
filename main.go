package main

import (
	"Lesson15/internal/controller"
	"Lesson15/internal/repository"
	"Lesson15/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.POST("/users", ctrl.Create)
	r.GET("/users/:id", ctrl.Get)
	r.PUT("/users/:id", ctrl.Update)
	r.DELETE("/users/:id", ctrl.Delete)

	fmt.Println("Server started at localhost:8080")
	r.Run(":8080")

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
