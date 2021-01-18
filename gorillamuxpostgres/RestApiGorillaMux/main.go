package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonesrohith2020/RestApiGorillaMux/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := app.App{}
	app.Initialize(
		os.Getenv("DB_Addr"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	app.Run(":8000")
}
