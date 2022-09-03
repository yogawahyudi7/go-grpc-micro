package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Id      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title   string
	Content string
	UserId  uuid.UUID `gorm:"type:uuid"`
}
