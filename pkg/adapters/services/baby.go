package services

import (
	"main/pkg/adapters/models"
	"main/pkg/adapters/repos"
)

type BabyService interface {
	getRandomBabies() []*models.Baby
}

type babyService struct {
	repo *repos.BabyRepo
}

func (service *babyService) getRandomBabies() []*models.Baby {
	return nil
}
