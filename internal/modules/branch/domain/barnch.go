package domain

import (
	
	"github.com/abdallahelassal/Store/internal/modules/user/domain"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model
	UUID uuid.UUID 	`json:"uuid" gorm:"default:uuid_generate_v4()"`
	Name string `json:"name"`
	User []domain.User `json:"user" gorm:"foreignKey:branchUUID;references:UUID"`
}

func (b *Branch) BeforeCreate(tx *gorm.DB){
	b.UUID = uuid.Must(uuid.NewV4())
}