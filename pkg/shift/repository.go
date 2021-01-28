package shift

import "gorm.io/gorm"

type Repository interface {
	// pass in any date and create the week containing that date if not exist
	Create(shift *Shift) (id uint, err error)
	Update(shift *Shift) error
	DeleteByID(id uint) error
	// input a date and return the week containing that date
	FindByID(id uint) (*Shift, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// pass in any date and create the week containing that date if not exist
func (r *repository) Create(shift *Shift) (id uint, err error) {
	result := r.db.Create(shift)
	if result.Error != nil {
		return 0, result.Error
	}

	return shift.ID, nil
}

func (r *repository) Update(shift *Shift) error {
	result := r.db.Save(shift)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) DeleteByID(id uint) error {
	result := r.db.Where("id = ?", id).Delete(&Shift{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// input a date and return the week containing that date
func (r *repository) FindByID(id uint) (*Shift, error) {
	panic("not implemented") // TODO: Implement
}
