package usecases

import (
	"github.com/hmrbcnto/prescription-api/entities"
	doctor_repo "github.com/hmrbcnto/prescription-api/infastructure/db/mongo/doctor_repo"
)

type DoctorUsecase interface {
	GetAllDoctors() ([]entities.Doctor, error)
	GetDoctorById(userId string) (*entities.Doctor, error)
	CreateDoctor(user *entities.Doctor) (*entities.Doctor, error)
}

type doctorUsecase struct {
	doctorRepo doctor_repo.DoctorRepo
}

func NewDoctorUsecase(doctorRepo doctor_repo.DoctorRepo) DoctorUsecase {
	return &doctorUsecase{
		doctorRepo: doctorRepo,
	}
}

func (doctorUsecase *doctorUsecase) GetAllDoctors() ([]entities.Doctor, error) {
	// Business logic if needed here

	users, err := doctorUsecase.doctorRepo.GetAllDoctors()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (doctorUsecase *doctorUsecase) GetDoctorById(userId string) (*entities.Doctor, error) {
	// More business logic here if needed

	user, err := doctorUsecase.doctorRepo.GetDoctorById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (doctorUsecase *doctorUsecase) CreateDoctor(user *entities.Doctor) (*entities.Doctor, error) {
	// Business logic again

	user, err := doctorUsecase.doctorRepo.CreateDoctor(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
