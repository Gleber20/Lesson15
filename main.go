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
	"github.com/rs/zerolog"
	"os"
)

// @title OnlineShop API
// @contact.name OnlineShop API Service
// @contact.url http://test.com
// @contact.email test@test.com
func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().Msg("Starting up application...")

	cfg := config.LoadConfig()
	db, err := sqlx.Open("postgres", cfg.DSN())
	if err != nil {
		logger.Error().Err(err).Msg("Error during connection to database")
		return
	}
	logger.Info().Msg("Read settings successfully")

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr(),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		logger.Error().Err(err).Msg("Ошибка подключения к Redis")
	}

	cache := repository.NewCache(rdb)
	logger.Info().Msg("Redis connected successfully")

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo, cache)
	ctrl := controller.NewUserController(svc)

	if err = ctrl.RunServer(":" + cfg.ServerPort); err != nil {
		logger.Error().Err(err).Msg("Error during running http-server")
	}

	if err = db.Close(); err != nil {
		logger.Error().Err(err).Msg("Error during closing database connection")
	}
}
