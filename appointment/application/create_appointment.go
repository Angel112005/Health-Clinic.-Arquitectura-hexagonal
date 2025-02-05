package appointment

import "my-clinic-api/appointment/domain"

type CreateAppointment struct {
	repo appointment.Repository
}

func NewCreateAppointment(repo appointment.Repository) *CreateAppointment {
	return &CreateAppointment{repo: repo}
}

func (uc *CreateAppointment) Execute(a appointment.Appointment) error {
	return uc.repo.Save(a)
}
