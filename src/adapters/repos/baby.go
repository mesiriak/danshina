package repos

import (
	"context"
	"log"
	"main/pkg/mongo"
	"main/src/adapters/models"
)

func CreateBabyRepo() *BabyRepo {
	return &BabyRepo{mongoClient: mongo.GetMongoClient()}
}

type BabyRepo struct {
	mongoClient mongo.IMongoClient
}

func (repo *BabyRepo) GetAllBabies() *models.BabyList {

	var response []*models.Baby

	cursor, count := repo.mongoClient.GetBulkObject("babies", 0, 0)

	err := cursor.All(context.TODO(), &response)

	if err != nil {
		log.Fatalf("Error happened during decoding\n%e", err)
	}

	return &models.BabyList{
		Count:   count,
		Results: response,
	}
}

func (repo *BabyRepo) GetRandomBabies() []*models.Baby {
	//TODO implement me
	panic("implement me")
}

func (repo *BabyRepo) CreateBaby() *models.Baby {
	//TODO implement me
	panic("implement me")
}

func (repo *BabyRepo) UpdateBabyCounter() bool {
	//TODO implement me
	panic("implement me")
}

func (repo *BabyRepo) DeleteBaby() bool {
	//TODO implement me
	panic("implement me")
}
