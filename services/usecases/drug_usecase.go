package usecases

import (
	"github.com/hmrbcnto/prescription-api/entities"
	"github.com/hmrbcnto/prescription-api/infastructure/db/mongo/drug_repo"
)

type DrugUsecase interface {
	GetAllDrugs() ([]entities.Drug, error)
	GetDrugById(string) (*entities.Drug, error)
	CreateDrug(*entities.Drug) (*entities.Drug, error)
	DeleteDrugById(string) (error)
	UpdateDrugById(string, *entities.Drug) (*entities.Drug, error)
}

type drugUsecase struct {
	drugRepo drug_repo.DrugRepo
}

func NewDrugUsecase(drugRepo drug_repo.DrugRepo) DrugUsecase {
	return &drugUsecase{
		drugRepo: drugRepo,
	}
}

func (drugUsecase *drugUsecase) GetAllDrugs() ([]entities.Drug, error) {
	// Insert business logic here
	drugs, err := drugUsecase.drugRepo.GetDrugs()

	if err != nil {
		return nil, err
	}

	return drugs, nil
}

func (drugUsecase *drugUsecase) GetDrugById(id string) (*entities.Drug, error) {
	// Inserted business logic here
	drug, err := drugUsecase.drugRepo.GetDrugById(id)

	if err != nil {
		return nil, err
	}

	return drug, nil
}

func (drugUsecase *drugUsecase) CreateDrug(drug *entities.Drug) (*entities.Drug, error) {
	// Insert business logic here

	drug, err := drugUsecase.drugRepo.CreateDrug(drug)

	if err != nil {
		return nil, err
	}

	return drug, nil
}

func (drugUsecase *drugUsecase) UpdateDrugById(id string, drug *entities.Drug) (*entities.Drug, error) {
	// Insert business logic here

	drug, err := drugUsecase.drugRepo.UpdateDrugById(id, drug)

	if err != nil {
		return nil, err
	}

	return drug, nil
}

func (drugUsecase *drugUsecase) DeleteDrugById(id string) error {
	// Business logic here
	err := drugUsecase.drugRepo.DeleteDrugById(id)

	return err;
}