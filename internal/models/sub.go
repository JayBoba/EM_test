package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	Id          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ServiceName string     `gorm:"not null" json:"service_name"`
	Price       int        `gorm:"not null;check:price >= 0" json:"price"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	StartDate   time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate     *time.Time `gorm:"type:date" json:"end_date,omitempty"`
}
