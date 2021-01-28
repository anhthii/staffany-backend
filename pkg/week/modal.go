package week

import (
	"time"

	"github.com/anhthii/staffany-backend/pkg/date"
	"gorm.io/gorm"
)

type Week struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	WeekNumber int `json:"week_number"`
	// store date as integer type for fast querying
	StartDate uint64 `json:"start_date"`

	IsPublished bool `json:"is_published"`
	UserID      uint `json:"user_id"`

	// has many relationship: A week has many dates
	Dates []date.Date `json:"dates"`
}
