package config

import (
	"os"
	"strconv"
)

type (
	Env struct {
		DB                     string
		Port                   string
		DatabaseContextTimeout int
	}
)

func Init() *Env {
	dbTimeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		dbTimeout = 10
	}
	return &Env{
		DB:                     os.Getenv("DB"),
		Port:                   ":" + os.Getenv("PORT"),
		DatabaseContextTimeout: dbTimeout,
	}
}
