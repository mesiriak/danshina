package producers

import "main/src/adapters/models"

type BabyProducer interface {
	GetAllBabies(limit int, offset int) *models.BabyListModel
	GetRandomBabies(size int) *models.BabyRandomListModel
	CreateBaby(model *models.BabyModel) *models.BabyModel
	UpdateBabyCounter(uuid string) *models.BabyModel
	DeleteBaby(uuid string) bool
}
