package producers

import "main/src/adapters/models"

type BabyProducer interface {
	GetAllBabies() *models.BabyList
	GetRandomBabies() []*models.Baby
	CreateBaby() *models.Baby
	UpdateBabyCounter(babyID rune) bool
	DeleteBaby(babyID rune) bool
}
