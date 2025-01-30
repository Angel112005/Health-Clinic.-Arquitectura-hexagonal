package doctor

type Repository interface {
	Save(Doctor) error
	FindAll() ([]Doctor, error)
	FindByID(id int) (*Doctor, error)
	Update(Doctor) error
	Delete(id int) error
}
