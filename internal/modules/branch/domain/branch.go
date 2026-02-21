package domain

import (
	
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model
	UUID 		uuid.UUID 			`json:"uuid" gorm:"type:uuid;uniqueIndex;"`
	Name 		string 				`json:"name" gorm:"not null"`
	
}

func (b *Branch) BeforeCreate(tx *gorm.DB) error {
	b.UUID = uuid.Must(uuid.NewV4())
	return nil
}

