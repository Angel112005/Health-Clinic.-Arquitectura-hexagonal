package appointment

import "my-clinic-api/appointment/domain"

type UpdateAppointment struct {
	repo appointment.Repository
}

func NewUpdateAppointment(repo appointment.Repository) *UpdateAppointment {
	return &UpdateAppointment{repo: repo}
}

func (uc *UpdateAppointment) Execute(a appointment.Appointment) error {
	return uc.repo.Update(a)
}
