package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() *EnvProperties {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	return &EnvProperties{
		DatabaseHost:   os.Getenv("DATABASE_HOST"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		DatabasePasswd: os.Getenv("DATABASE_PASSWD"),
		DatabaseUser:   os.Getenv("DATABASE_USER"),
		DatabasePort:   os.Getenv("DATABASE_PORT"),
	}

}
