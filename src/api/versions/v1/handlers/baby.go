package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"main/src/api/dependencies"
	"strconv"
)

func GetAllBabies(ctx *fiber.Ctx) error {
	response := dependencies.ServerManagerInstance.BabyService.GetAllBabies()

	return ctx.JSON(response)
}

func GetRandomBabies(ctx *fiber.Ctx) error {
	response := dependencies.ServerManagerInstance.BabyService.GetRandomBabies()

	return ctx.JSON(response)
}

func CreateBaby(ctx *fiber.Ctx) error {
	// TODO: make cdn uploading here
	//pictureUrl := dependencies.ServerManagerInstance.PictureService
	response := dependencies.ServerManagerInstance.BabyService.CreateBaby()

	return ctx.JSON(response)
}

func UpdateBabyCounter(ctx *fiber.Ctx) error {
	babyID, err := strconv.ParseInt(ctx.Params("id"), 10, 32)

	if err != nil {
		log.Println("Error happened during request processing")
	}

	response := dependencies.ServerManagerInstance.BabyService.UpdateBabyCounter(rune(babyID))

	return ctx.JSON(response)
}

func DeleteBaby(ctx *fiber.Ctx) error {
	babyID, err := strconv.ParseInt(ctx.Params("id"), 10, 32)

	if err != nil {
		log.Println("Error happened during request processing")
	}

	response := dependencies.ServerManagerInstance.BabyService.DeleteBaby(rune(babyID))

	return ctx.JSON(response)
}
