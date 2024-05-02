package repository

import (
	"context"
	"github.com/b3liv3r/subscription-for-gym/modules/subscription/models"
)

type SubscriberRepository interface {
	Create(ctx context.Context, userID int) error
	GetByID(ctx context.Context, userID int) (models.Subscription, error)
	UpdateType(ctx context.Context, userID int, subsType int, monthCount int) error
	ExtendEndDate(ctx context.Context, userID int, monthCount int) error
}
