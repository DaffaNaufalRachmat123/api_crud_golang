package models

import "time"

type User struct {
	ID			uint		`gorm:"primary_key" json:"id"`
	Username	string		`json:"username"; gorm:"type:varchar(255);NOT NULL"`
	Password	string		`json:"password"; gorm:"type:varchar(255);NOT NULL"`
	Email		string		`json:"email"; gorm:"type:varchar(255);NOT NULL"`
	Unique_Key	string		`json:"unique_Key"; gorm:"type:varchar(255);NOT NULL"`
	Profile		Profile		`json:"profile"; gorm:"foreignkey:user_id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updatead_at"`
}

type LoginValidate struct {
	Username	string	`json:"username" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

type RegisterValidate struct {
	Username	string	`json:"username" binding:"required"`
	Password	string	`json:"password" binding:"required"`
	Email		string	`json:"email" binding:"required"`
	Name		string	`json:"name" binding:"required"`
	Phone		string	`json:"phone" binding:"required"`
	City		string	`json:"city" binding:"required"`
}