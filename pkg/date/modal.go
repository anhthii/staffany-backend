package date

import (
	"github.com/anhthii/anystaff-backend/pkg/shift"
	"gorm.io/gorm"
)

type Date struct {
	gorm.Model

	Date uint

	WeekID uint
	UserID uint

	IsPublished bool

	// has many relationship: a date has many shifts
	Shifts []shift.Shift
}
