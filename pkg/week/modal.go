package week

import (
	"github.com/anhthii/anystaff-backend/pkg/date"
	"gorm.io/gorm"
)

type Week struct {
	gorm.Model

	// store date as integer type for fast querying
	StartDate uint64
	EndDate   uint64

	IsPublished bool
	UserID      uint

	// has many relationship: A week has many dates
	Dates []date.Date
}
