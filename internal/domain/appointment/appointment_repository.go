package appointment

type Repository interface {
	Save(Appointment) error
	FindAll() ([]Appointment, error)
	FindByID(id int) (*Appointment, error)
	Update(Appointment) error
	Delete(id int) error
}
