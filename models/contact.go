package models

import (
	"time"
)

type Contact struct {

	ID uint `gorm:"primarykey;column:id;type:BIGINT UNSINED AUTO_INCREMENT"`

	FullName string `gorm:"column:full_name;type:VARCHAR(10);not null"`

	Email string `gorm:"colum:email;type:VARCHAR(100);not null"`

	Phone string `gorm:"colum:phone;type:VARCHAR(15);not null"`

	Message string `gorm:"colum:message;type:TEXT;not null"`

	CreatedAt time.Time `gorm:"column:createdAt;type:DATETIME;autoCreateTime"`

	UpdatedAt time.Time `gorm:"column:updatedAt;type:DATETIME;autoUpdateTime"`

	DeletedAt time.Time `gorm:"column:deletedAt;type:DATETIME;index"`
	
}

func (Contact) TableName() string {
	return "contacts_messages"
}