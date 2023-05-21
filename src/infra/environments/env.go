package environments

import (
	"log"

	"github.com/joho/godotenv"
)

type Environment struct {
	DATABASE_URL string
}

var datasetEnv *Environment

func GetEnv() *Environment {
	return datasetEnv
}

func StartEnv() {
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}
	datasetEnv = &Environment{
		DATABASE_URL: envs["DATABASE_URL"],
	}
}
