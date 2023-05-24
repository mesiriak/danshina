package main

import (
	"fmt"
	"log"
	"main/cfg"
	"main/src"
)

func main() {
	config := cfg.GetConfig(".")

	app := src.CreateApp(config)

	appAddr := fmt.Sprintf(":%d", config.AppPort)

	log.Fatal(app.Listen(appAddr))
}
