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

type EmployeeService struct {
	repo  contracts.RepositoryI
	cache contracts.CacheI
}

func NewUserService(repo contracts.RepositoryI, cache *repository.Cache) *EmployeeService {
	return &EmployeeService{
		repo:  repo,
		cache: cache,
	}
}

var (
	defaultTTL = time.Minute * 5
)

func (s *EmployeeService) CreateEmployee(ctx context.Context, user models.Employee) error {
	err := s.repo.CreateEmployee(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, userId int) (models.Employee, error) {
	var user models.Employee

	// 1. Тут берем из кеша данные
	cacheKey := fmt.Sprintf("user_%d", userId)
	err := s.cache.Get(ctx, cacheKey, &user)
	if err == nil {
		return user, nil
	}
	// Если ошибка не связана с отсутствием ключа — возвращаем её
	if !errors.Is(err, redis.Nil) {
		return models.Employee{}, err
	}

	// 2. Если в кэше нет, обращаемся в репозиторий, да
	user, err = s.repo.GetEmployeeByID(ctx, userId)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return models.Employee{}, errs.ErrUserNotFound
		}
		return models.Employee{}, err
	}

	// 3. Сохраняем в кэш на 5 минут
	if err := s.cache.Set(ctx, cacheKey, user, defaultTTL); err != nil {
		log.Printf("failed to set user cache: %v", err)
	}

	return user, nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, user models.Employee) error {
	_, err := s.repo.GetEmployeeByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotFound
		}
		return err
	}
	return nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, userId int) error {
	_, err := s.repo.GetEmployeeByID(ctx, userId)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotFound
		}
		return err
	}
	err = s.repo.DeleteEmployee(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}
