package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
) 

type User struct{
	gorm.Model
	UUID		uuid.UUID	`json:"uuid" gorm:"default:uuid_generate_v4()"`
	User_name	string 		`json:"username" validate:"required,min=6"`
	Email		string		`json:"email" validate:"required,email"`
	Password	string		`json:"password" validate:"required,password"`
	Role		Role		`json:"role" gorm:"default:user" validate:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB){
	u.UUID = uuid.Must(uuid.NewV4())
}