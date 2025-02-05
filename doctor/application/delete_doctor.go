package doctor

import "my-clinic-api/doctor/domain"

type DeleteDoctor struct {
	repo doctor.Repository
}

func NewDeleteDoctor(repo doctor.Repository) *DeleteDoctor {
	return &DeleteDoctor{repo: repo}
}

func (uc *DeleteDoctor) Execute(id int) error {
	return uc.repo.Delete(id)
}
