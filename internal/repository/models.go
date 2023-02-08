package repository

import "time"

type TestTable struct {
	ID    int64  `gorm:"column:id;primary_key;autoIncrement"`
	Title string `gorm:"column:title"`
}

type Users struct {
	ID          string    ` json:"id" gorm:"id"`
	Login       string    `json:"login" gorm:"login"`
	Password    string    `json:"password" gorm:"password"`
	Name        string    `json:"name" gorm:"name"`
	AccessToken string    `json:"access_token" gorm:"column:access_token"`
	Active      bool      `json:"active" gorm:"active"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type Tokens struct {
	UserID    string    ` json:"user_id" gorm:"column:user_id"`
	Token     string    `json:"token" gorm:"token"`
	Platform  string    `json:"platform" gorm:"platform"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
