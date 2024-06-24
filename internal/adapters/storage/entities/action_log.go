package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ActionLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    *uint
	Action    string
	IP        string
	Endpoint  string
	Payload   map[string]any `gorm:"serializer:json"`
	Method    string
}

// BeforeCreate is a GORM hook that is called before a new record is inserted into the database
func (log *ActionLog) BeforeCreate(tx *gorm.DB) (err error) {
	log.ID = uuid.New()
	return
}
