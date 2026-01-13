package model

import "time"

type Admin struct {
	AdminID      int       `gorm:"column:adminid;primaryKey;autoIncrement" json:"admin_id"`
	Username     string    `gorm:"column:username;type:varchar(100);not null;unique" json:"username"`
	PasswordHash string    `gorm:"column:passwordhash;type:varchar(255);not null" json:"-"`
	FullName     string    `gorm:"column:fullname;type:varchar(150);not null" json:"full_name"`
	Email        string    `gorm:"column:email;type:varchar(150);unique" json:"email"`
	Phone        string    `gorm:"column:phone;type:varchar(20)" json:"phone"`
	IsActive     bool      `gorm:"column:isactive;not null;default:true" json:"is_active"`
	CreatedAt    time.Time `gorm:"column:createdtime;autoCreateTime" json:"created_at"`
}

func (Admin) TableName() string {
	return "admin"
}
