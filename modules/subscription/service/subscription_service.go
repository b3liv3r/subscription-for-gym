package service

import (
	"context"
	"github.com/b3liv3r/subscription-for-gym/modules/subscription/models"
	"github.com/b3liv3r/subscription-for-gym/modules/subscription/repository"
	"go.uber.org/zap"
)

type SubscriptionService struct {
	repo repository.SubscriberRepository
	log  *zap.Logger
}

func NewSubscriptionService(repo repository.SubscriberRepository, log *zap.Logger) Subscriber {
	return &SubscriptionService{repo: repo, log: log}
}

func (s *SubscriptionService) Create(ctx context.Context, userID int) (string, error) {
	err := s.repo.Create(ctx, userID)
	if err != nil {
		s.log.Error("Ошибка при создании подписки", zap.Error(err))
		return "", err
	}
	return "Подписка успешно создана", nil
}

func (s *SubscriptionService) GetByID(ctx context.Context, userID int) (models.Subscription, error) {
	subscription, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		s.log.Error("Ошибка при получении подписки", zap.Error(err))
		return models.Subscription{}, err
	}
	return subscription, nil
}

func (s *SubscriptionService) UpdateType(ctx context.Context, userID int, subsType int, monthCount int) (string, error) {
	err := s.repo.UpdateType(ctx, userID, subsType, monthCount)
	if err != nil {
		s.log.Error("Ошибка при обновлении типа подписки", zap.Error(err))
		return "", err
	}
	return "Тип подписки успешно обновлен", nil
}

func (s *SubscriptionService) ExtendEndDate(ctx context.Context, userID int, monthCount int) (string, error) {
	err := s.repo.ExtendEndDate(ctx, userID, monthCount)
	if err != nil {
		s.log.Error("Ошибка при продлении даты окончания подписки", zap.Error(err))
		return "", err
	}
	return "Дата окончания подписки успешно продлена", nil
}
