package repos

import (
	"context"
	"log"
	"main/pkg/mongo"
	"main/src/adapters/models"
)

func CreateBabyRepo() *BabyRepo {
	return &BabyRepo{mongoClient: mongo.GetMongoClient(), collection: "babies"}
}

type BabyRepo struct {
	mongoClient mongo.IMongoClient
	collection  string
}

func (repo *BabyRepo) GetAllBabies(limit int, offset int) *models.BabyListModel {
	var results []*models.BabyModel

	cursor, count := repo.mongoClient.GetBulkObject(repo.collection, limit, offset)

	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatalf("Error happened during decoding\n%e", err)
	}

	return &models.BabyListModel{
		Count:   count,
		Results: results,
	}
}

func (repo *BabyRepo) GetRandomBabies(size int) *models.BabyRandomListModel {
	var results []*models.BabyModel

	cursor, returnSize := repo.mongoClient.GetRandomObjects(repo.collection, size)

	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatalf("Error happened during decoding\n%e", err)
	}

	return &models.BabyRandomListModel{
		Size:    returnSize,
		Results: results,
	}
}

func (repo *BabyRepo) CreateBaby(model *models.BabyModel) *models.BabyModel {
	var results *models.BabyModel

	cursor := repo.mongoClient.CreateObject(repo.collection, model)

	if err := cursor.Decode(&results); err != nil {
		log.Fatalf("Error happened during decoding\n%e", err)
	}

	return results
}

func (repo *BabyRepo) UpdateBabyCounter(uuid string) *models.BabyModel {
	var results *models.BabyModel

	cursor := repo.mongoClient.UpdateObjectCounter(repo.collection, uuid)

	if err := cursor.Decode(&results); err != nil {
		log.Fatalf("Error happened during decoding\n%e", err)
	}

	return results
}

func (repo *BabyRepo) DeleteBaby(uuid string) bool {
	//TODO implement me
	panic("implement me")
}
