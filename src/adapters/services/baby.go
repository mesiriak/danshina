package services

import (
	"main/src/adapters/models"
	"main/src/adapters/repos"
)

func CreateBabyService() *BabyService {
	return &BabyService{repo: repos.CreateBabyRepo()}
}

type BabyService struct {
	repo *repos.BabyRepo
}

func (service *BabyService) GetAllBabies(limit int, offset int) *models.BabyListModel {

	return service.repo.GetAllBabies(limit, offset)
}

func (service *BabyService) GetRandomBabies(size int) *models.BabyRandomListModel {

	return service.repo.GetRandomBabies(size)
}

func (service *BabyService) CreateBaby(model *models.BabyModel) *models.BabyModel {

	return service.repo.CreateBaby(model)
}

func (service *BabyService) UpdateBabyCounter(uuid string) *models.BabyModel {

	return service.repo.UpdateBabyCounter(uuid)
}

func (service *BabyService) DeleteBaby(uuid string) bool {
	//TODO implement me
	panic("implement me")
}
