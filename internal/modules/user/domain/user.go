package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
) 

type User struct{
	
	UUID		uuid.UUID		`json:"uuid" gorm:"type:uuid;not null;uniqueIndex"`
	Username	string 			`json:"username" validate:"required,min=6"`
	Email		string			`json:"email" validate:"required,email"`
	Password	string			`json:"-" validate:"required,password"`
	Role		Role			`json:"role" validate:"required" gorm:"type:varchar(20)"`
	BranchUUID  *uuid.UUID		`json:"branch_uuid"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

}

func (u *User) BeforeCreate(tx *gorm.DB)error{
	u.UUID = uuid.Must(uuid.NewV4())
	return nil
}