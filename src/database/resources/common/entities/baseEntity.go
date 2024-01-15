package entities

import (
	"github.com/google/uuid"
	"time"
)

type BaseEntity struct {
	ID        uuid.UUID `gorm:"type:uuid; primary_key; unique; default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
