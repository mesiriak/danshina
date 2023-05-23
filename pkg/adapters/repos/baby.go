package repos

import "main/pkg/adapters/models"

type BabyRepo interface {
	getRandomBabies() []*models.Baby
}

type babyRepo struct {
}

func (repo *babyRepo) getRandomBabies() []*models.Baby {
	return nil
}
