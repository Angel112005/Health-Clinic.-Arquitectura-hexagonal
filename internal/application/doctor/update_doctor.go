package doctor

import "my-clinic-api/internal/domain/doctor"

type UpdateDoctor struct {
	repo doctor.Repository
}

func NewUpdateDoctor(repo doctor.Repository) *UpdateDoctor {
	return &UpdateDoctor{repo: repo}
}

func (uc *UpdateDoctor) Execute(d doctor.Doctor) error {
	return uc.repo.Update(d)
}
