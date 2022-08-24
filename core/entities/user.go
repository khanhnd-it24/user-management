package entities

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey,auto_increment"`
	Username  string    `json:"username" gorm:"unique;notnull"`
	Password  []byte    `json:"-" gorm:"notnull"`
	Role      string    `json:"role" gorm:"notnull"`
	Active    bool      `json:"active" gorm:"default:true"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Users []*User
