package entity

import "time"

type PasswordItem struct {
	ID          int       `gorm:"primaryKey" json:"password_id"`
	Website     string    `json:"website"`
	WebPassword string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
}

type PasswordItemInput struct {
	ID          int    `gorm:"primaryKey" json:"password_id"`
	Website     string `json:"website"`
	WebPassword string `json:"password"`
}

type PasswordItemUpdateInput struct {
	Website     string `json:"website"`
	WebPassword string `json:"password"`
}
