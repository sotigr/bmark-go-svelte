package system

import "os"

const (
	ENV_DEV  = iota
	ENV_PROD = iota
)

func GetEnv() int {
	envStr := os.Getenv("ENV")
	if envStr == "dev" {
		return ENV_DEV
	} else if envStr == "prod" {
		return ENV_PROD
	} else {
		panic("Fatal error the ENV environmental variable is either not provided or not supported.")
	}
}
