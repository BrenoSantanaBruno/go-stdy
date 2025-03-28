package config

import (
	"log"
	"os"
)

type Config struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func New() *Config {
	return &Config{
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
