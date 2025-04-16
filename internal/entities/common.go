package entities

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (s *Base) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("Before Create di panggil !")
	s.ID = uuid.New().String()
	return
}
