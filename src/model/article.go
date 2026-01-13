package model

import "time"

type Article struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Title    string    `json:"title" gorm:"type:text;not null"`
	Content  string    `json:"content" gorm:"type:text;not null"`
	Category string    `json:"category" gorm:"type:varchar(255);not null;index"`
	Author   string    `json:"author" gorm:"type:varchar(255);not null"`
	Image    *string   `json:"image"`
	Created   time.Time `json:"created" gorm:"column:created_at;index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;index"`
	Featured  bool      `json:"featured" gorm:"default:false;index"`
}

type Response struct {
	Message string      `json:"message"`
	Article *Article    `json:"article,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ImageUploadResponse struct {
	Message string `json:"message"`
	URL     string `json:"url"`
}
