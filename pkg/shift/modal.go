package shift

import "gorm.io/gorm"

type Shift struct {
	gorm.Model

	DateID uint
	UserID uint

	QuarterStart uint
	NumQuarter   uint

	Title       string
	Description string
}
