package service

import (
	"Lesson15/internal/contracts"
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"Lesson15/internal/repository"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type UserService struct {
	repo  contracts.RepositoryI
	cache contracts.CacheI
}

func NewUserService(repo contracts.RepositoryI, cache *repository.Cache) *UserService {
	return &UserService{
		repo:  repo,
		cache: cache,
	}
}

var (
	defaultTTL = time.Minute * 5
)

func (s *UserService) CreateUser(ctx context.Context, user models.User) error {
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserById(ctx context.Context, userId int) (models.User, error) {
	var user models.User

	// 1. Тут берем из кеша данные
	cacheKey := fmt.Sprintf("user_%d", userId)
	err := s.cache.Get(ctx, cacheKey, &user)
	if err == nil {
		return user, nil
	}
	// Если ошибка не связана с отсутствием ключа — возвращаем её
	if !errors.Is(err, redis.Nil) {
		return models.User{}, err
	}

	// 2. Если в кэше нет, обращаемся в репозиторий, да
	user, err = s.repo.GetUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return models.User{}, errs.ErrUserNotfound
		}
		return models.User{}, err
	}

	// 3. Сохраняем в кэш на 5 минут
	if err := s.cache.Set(ctx, cacheKey, user, defaultTTL); err != nil {
		log.Printf("failed to set user cache: %v", err)
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user models.User) error {
	_, err := s.repo.GetUserById(ctx, user.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotfound
		}
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, userId int) error {
	_, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotfound
		}
		return err
	}
	err = s.repo.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}
