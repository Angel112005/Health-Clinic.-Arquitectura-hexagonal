package doctor

import "my-clinic-api/internal/domain/doctor"

type ListDoctors struct {
	repo doctor.Repository
}

func NewListDoctors(repo doctor.Repository) *ListDoctors {
	return &ListDoctors{repo: repo}
}

func (uc *ListDoctors) Execute() ([]doctor.Doctor, error) {
	return uc.repo.FindAll()
}
