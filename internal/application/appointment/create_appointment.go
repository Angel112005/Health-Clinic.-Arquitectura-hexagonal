package appointment

import "my-clinic-api/internal/domain/appointment"

type CreateAppointment struct {
	repo appointment.Repository
}

func NewCreateAppointment(repo appointment.Repository) *CreateAppointment {
	return &CreateAppointment{repo: repo}
}

func (uc *CreateAppointment) Execute(a appointment.Appointment) error {
	return uc.repo.Save(a)
}
