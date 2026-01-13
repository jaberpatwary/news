package model

import (
	"time"
)

type Comment struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int       `gorm:"not null" json:"user_id"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	IsAnonymous bool      `gorm:"not null;default:false" json:"is_anonymous"`
	IsDeleted   bool      `gorm:"not null;default:false" json:"is_deleted"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Association / Foreign Key
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}
