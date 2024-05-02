package service

import (
	"context"
	"github.com/b3liv3r/subscription-for-gym/modules/subscription/models"
)

type Subscriber interface {
	Create(ctx context.Context, userId int) (string, error)
	GetByID(ctx context.Context, userId int) (models.Subscription, error)
	UpdateType(ctx context.Context, userID int, subsType int, monthCount int) (string, error)
	ExtendEndDate(ctx context.Context, userID int, monthCount int) (string, error)
}
