package main

import (
	"backend/internal/app"
	"log"
)

func main() {
	app := app.New()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
