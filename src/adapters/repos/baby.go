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

func (repo *BabyRepo) GetAllBabies(limit int, offset int) *models.BabyListModel {
	var results []*models.BabyModel

	cursor, count := repo.mongoClient.GetBulkObject("babies", limit, offset)

	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatalf("Error happened during decoding\n%e", err)
	}

	return &models.BabyListModel{
		Count:   count,
		Results: results,
	}
}

func (repo *BabyRepo) GetRandomBabies() []*models.BabyModel {
	//TODO implement me
	panic("implement me")
}

func (repo *BabyRepo) CreateBaby(model *models.BabyModel) *models.BabyModel {
	var results *models.BabyModel

	cursor := repo.mongoClient.CreateObject("babies", model)

	if err := cursor.Decode(&results); err != nil {
		log.Fatalf("Error happened during decoding\n%e", err)
	}

	return results
}

func (repo *BabyRepo) UpdateBabyCounter(uuid string) *models.BabyModel {
	var results *models.BabyModel

	cursor := repo.mongoClient.UpdateObject("babies", uuid)

	if err := cursor.Decode(&results); err != nil {
		log.Fatalf("Error happened during decoding\n%e", err)
	}

	return results
}

func (repo *BabyRepo) DeleteBaby(uuid string) bool {
	//TODO implement me
	panic("implement me")
}
