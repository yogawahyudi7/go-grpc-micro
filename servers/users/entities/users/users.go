package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string
	Image_Url string
}
