package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	valString := os.Getenv(key)

	valInt, err := strconv.Atoi(valString)
	if err != nil {
		return fallback
	}

	return valInt
}
