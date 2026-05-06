package main

import (
	"slskd-exporter/extract"
)

func main() {
	env := GetEnvMap()
	extract.Extract(env.Slskd)
}
