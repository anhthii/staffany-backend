package date

import "gorm.io/gorm"

type Repository interface {
	// pass in any date and create the week containing that date if not exist
	Create(date *Date) (id uint, err error)
	Publish() error
}

type repository struct {
	db *gorm.DB
}

// pass in any date and create the week containing that date if not exist
func (r *repository) Create(date *Date) (id uint, err error) {

	result := r.db.Create(date)
	if result.Error != nil {
		return 0, result.Error
	}

	return date.ID, nil
}

func (r *repository) Publish() error {
	panic("not implemented") // TODO: Implement
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
