package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}

	env, ok := envMap[key]
	if !ok {
		panic(fmt.Sprintf("key %s was not found", key))
	}

	return env
}
