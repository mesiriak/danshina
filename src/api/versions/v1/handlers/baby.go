package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"main/src/adapters/models"
	"main/src/api/dependencies"
	"strconv"
)

func GetAllBabies(ctx *fiber.Ctx) error {
	// TODO: if statements parsing can be replaced onto QueryModel -> ctx.ParamsParser(out QueryModel{})
	queryLimit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		log.Fatalf("Error to parse limit parameter.\n%e", err)
	}
	queryOffset, err := strconv.Atoi(ctx.Query("offset", "0"))
	if err != nil {
		log.Fatalf("Error to parse limit parameter.\n%e", err)
	}

	response := dependencies.ServerManagerInstance.BabyService.GetAllBabies(queryLimit, queryOffset)

	return ctx.JSON(response)
}

func GetRandomBabies(ctx *fiber.Ctx) error {
	querySize, err := strconv.Atoi(ctx.Query("size", "2"))

	if err != nil {
		log.Fatalf("Error to parse size parameter.\n%e", err)
	}

	response := dependencies.ServerManagerInstance.BabyService.GetRandomBabies(querySize)

	return ctx.JSON(response)
}

func CreateBaby(ctx *fiber.Ctx) error {
	var payload = models.CreateBabyModel{}

	if err := ctx.BodyParser(&payload); err != nil {
		log.Fatalf("Error happened during body validation.\n%e", err)
	}

	// TODO: make cdn uploading here
	//pictureUrl := dependencies.ServerManagerInstance.PictureService
	pictureUrl := "somedummyurl"

	baby := models.BabyModel{
		Uuid:       uuid.NewString(),
		Nickname:   payload.Nickname,
		PictureUrl: pictureUrl,
	}

	response := dependencies.ServerManagerInstance.BabyService.CreateBaby(&baby)

	return ctx.JSON(response)
}

func UpdateBabyCounter(ctx *fiber.Ctx) error {
	babyUuid := ctx.Params("uuid")

	response := dependencies.ServerManagerInstance.BabyService.UpdateBabyCounter(babyUuid)

	return ctx.JSON(response)
}

func DeleteBaby(ctx *fiber.Ctx) error {
	babyUuid := ctx.Params("uuid")

	response := dependencies.ServerManagerInstance.BabyService.DeleteBaby(babyUuid)

	return ctx.JSON(response)
}
