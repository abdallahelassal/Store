package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	
	UUID 			uuid.UUID 		`json:"uuid" gorm:"type:uuid;primaryKey"`
	Name			string			`json:"name" validate:"required,min=3,max=100"`
	MobileNumber 	string			`json:"mobile_number" gorm:"uniqueIndex" validate:"required,len=11"`
	BranchUUID  	*uuid.UUID 		`gorm:"type:uuid"`
	
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt	`gorm:"index"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	c.UUID = uuid.Must(uuid.NewV4())
	return nil
}