package model

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Email          string    `json:"email" gorm:"unique"`
	Password       string    `json:"password"`
	LineId         string    `json:"line_id"`
	IsSubscribed   string    `json:"is_subscribed"`
	StripeCustomer string    `json:"stripe_customer"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
