package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//utilizamos la librería godotenv, con esto recorremos las variables del archivo .env y las setean en el environment.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}