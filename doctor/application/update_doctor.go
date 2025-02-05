package doctor

import "my-clinic-api/doctor/domain"

type UpdateDoctor struct {
	repo doctor.Repository
}

func NewUpdateDoctor(repo doctor.Repository) *UpdateDoctor {
	return &UpdateDoctor{repo: repo}
}

func (uc *UpdateDoctor) Execute(d doctor.Doctor) error {
	return uc.repo.Update(d)
}
