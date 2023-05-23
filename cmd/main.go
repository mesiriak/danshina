package main

import (
	"fmt"
	"log"
	"main/pkg"
)

func main() {
	config := GetConfig()

	app := pkg.CreateApp(config)

	appAddr := fmt.Sprintf(":%d", config.AppPort)

	log.Fatal(app.Listen(appAddr))
}
