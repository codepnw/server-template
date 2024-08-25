package helpers

import (
	"os"
	"strconv"
)

func GetEnv(key, value string) string {
	if v, exists := os.LookupEnv(key); exists {
		return v
	}
	return value
}

func GetEnvAsBool(name string) bool {
	val := GetEnv(name, "")
	if res, err := strconv.ParseBool(val); err != nil {
		return res
	}
	return false
}
