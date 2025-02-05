package doctor

import (
	domain "my-clinic-api/doctor/domain"
)

type ListDoctorByID struct {
	repo domain.Repository
}

func NewListDoctorByID(repo domain.Repository) *ListDoctorByID {
	return &ListDoctorByID{repo: repo}
}

func (uc *ListDoctorByID) Execute(id int) (*domain.Doctor, error) {
	return uc.repo.FindByID(id)
}
