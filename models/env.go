package models

import "time"

type SlskdEnv struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	CERT     string
}

type DbEnv struct {
	USER     string
	PASSWORD string
	DB       string
	HOST     string
	PORT     string
}

type Env struct {
	Slskd          SlskdEnv
	DbEnv          DbEnv
	LogLevel       string
	ScrapeInterval time.Duration
}
