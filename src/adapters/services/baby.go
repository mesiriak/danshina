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

func (service *BabyService) GetAllBabies() *models.BabyList {

	return service.repo.GetAllBabies()
}

func (service *BabyService) GetRandomBabies() []*models.Baby {
	//TODO implement me
	panic("implement me")
}

func (service *BabyService) CreateBaby() *models.Baby {
	//TODO implement me
	panic("implement me")
}

func (service *BabyService) UpdateBabyCounter(babyID rune) bool {
	//TODO implement me
	panic("implement me")
}

func (service *BabyService) DeleteBaby(babyID rune) bool {
	//TODO implement me
	panic("implement me")
}
