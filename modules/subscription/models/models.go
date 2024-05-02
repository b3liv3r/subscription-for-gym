package models

import "time"

type Subscription struct {
	Id               int       `db:"user_id"`
	SubscriptionType int       `db:"subscription_type"`
	StartDate        time.Time `db:"start_date"`
	EndDate          time.Time `db:"end_date"`
}
