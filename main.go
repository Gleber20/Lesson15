package main

import (
	"Lesson15/internal/config"
	"Lesson15/internal/controller"
	"Lesson15/internal/repository"
	"Lesson15/internal/service"
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr(),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Ошибка подключения к Redis:", err)
	}

	cache := repository.NewCache(rdb)

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo, cache)
	ctrl := controller.NewUserController(svc)

	if err = ctrl.RunServer(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
