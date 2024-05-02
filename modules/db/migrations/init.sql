CREATE TABLE IF NOT EXISTS user_subscriptions (
    user_id INT UNIQUE,
    subscription_type INT,
    start_date DATE,
    end_date DATE
);
