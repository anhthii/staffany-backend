package date

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	// pass in any date and create the week containing that date if not exist
	Create(date *Date) (*Date, error)
	Publish() error
}

type repository struct {
	db *gorm.DB
}

// pass in any date and create the week containing that date if not exist
func (r *repository) Create(date *Date) (*Date, error) {
	err := r.db.Debug().Where("date = ?", date.Date).First(&date).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		result := r.db.Create(date)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return date, nil
}

func (r *repository) Publish() error {
	panic("not implemented") // TODO: Implement
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
