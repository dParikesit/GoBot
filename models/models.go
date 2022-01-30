package models

import (
	"time"
)

type Reminder struct {
	ID          uint `gorm:"primaryKey"`
	Schedule    time.Time
	Description string
	Done        bool
	CreatedAt   time.Time
}
