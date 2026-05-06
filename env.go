package main

import (
	"os"
	"strings"

	"slskd-exporter/domain"

	"github.com/joho/godotenv"
)

func GetEnvMap() domain.Env {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	env := make(map[string]string)

	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		env[parts[0]] = parts[1]
	}

	slskdEnv := domain.SlskdEnv{
		HOST:            env["SLSKD_HOST"],
		PORT:            env["SLSKD_PORT"],
		USER:            env["SLSKD_USER"],
		PASSWORD:        env["SLSKD_PASSWORD"],
		SCRAPE_INTERVAL: env["SCRAPE_INTERVAL"],
	}

	exposerEnv := domain.ExposerEnv{
		PORT: env["PORT"],
	}

	return domain.Env{
		Slskd:   slskdEnv,
		Exposer: exposerEnv,
	}
}
