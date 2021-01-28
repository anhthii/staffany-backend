package date

import (
	"time"

	"github.com/anhthii/anystaff-backend/pkg/shift"
	"gorm.io/gorm"
)

type Date struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Date uint64 `json:"date"`

	WeekID uint `json:"week_id"`
	UserID uint `json:"user_id"`

	IsPublished bool `json:"is_published"`

	// has many relationship: a date has many shifts
	Shifts []shift.Shift `json:"shifts"`
}
