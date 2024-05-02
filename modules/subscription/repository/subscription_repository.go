package repository

import (
	"context"
	"github.com/b3liv3r/subscription-for-gym/modules/subscription/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type SubscriptionRepositoryDB struct {
	db *sqlx.DB
}

func NewSubscriptionRepositoryDB(db *sqlx.DB) SubscriberRepository {
	return &SubscriptionRepositoryDB{
		db: db,
	}
}

func (r *SubscriptionRepositoryDB) Create(ctx context.Context, userID int) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO user_subscriptions (user_id, subscription_type, start_date, end_date) VALUES ($1, $2, $3, $4)", userID, 0, time.Now(), time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepositoryDB) GetByID(ctx context.Context, userID int) (models.Subscription, error) {
	var subscription models.Subscription
	err := r.db.GetContext(ctx, &subscription, "SELECT * FROM user_subscriptions WHERE user_id = $1", userID)
	if err != nil {
		return models.Subscription{}, err
	}
	return subscription, nil
}

func (r *SubscriptionRepositoryDB) UpdateType(ctx context.Context, userID int, Type int, monthCount int) error {
	currentDate := time.Now().Format("2006-01-02")
	endDate := time.Now().AddDate(0, monthCount, 0).Format("2006-01-02")

	_, err := r.db.ExecContext(ctx, "UPDATE user_subscriptions SET subscription_type = $1, start_date = $2, end_date = $3 WHERE user_id = $4", Type, currentDate, endDate, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepositoryDB) ExtendEndDate(ctx context.Context, userID int, monthCount int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE user_subscriptions SET end_date = end_date + INTERVAL '1 month' * $1 WHERE user_id = $2", monthCount, userID)
	if err != nil {
		return err
	}
	return nil
}
