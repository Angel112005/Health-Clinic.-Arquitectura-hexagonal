package appointment

import "my-clinic-api/internal/domain/appointment"

type ListAppointments struct {
	repo appointment.Repository
}

func NewListAppointments(repo appointment.Repository) *ListAppointments {
	return &ListAppointments{repo: repo}
}

func (uc *ListAppointments) Execute() ([]appointment.Appointment, error) {
	return uc.repo.FindAll()
}
