package appointment

import "my-clinic-api/internal/domain/appointment"

type DeleteAppointment struct {
	repo appointment.Repository
}

func NewDeleteAppointment(repo appointment.Repository) *DeleteAppointment {
	return &DeleteAppointment{repo: repo}
}

func (uc *DeleteAppointment) Execute(id int) error {
	return uc.repo.Delete(id)
}
