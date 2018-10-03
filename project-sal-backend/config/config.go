package config

import (
	"log"
	"os"
	"strings"
)

var (
	configs      map[string]string
	validConfigs map[string]bool
)

func init() {
	validConfigs = map[string]bool{
		"DB_HOST":     true,
		"DB_USER":     true,
		"DB_PASSWORD": true,
		"DB_PORT":     true,
		"DB_NAME":     true,
	}
}

func parseConfig() {
	configs = make(map[string]string)
	log.Printf("ENVS: %#v", os.Environ())
	for _, env := range os.Environ() {
		e := strings.Split(env, "=")
		configs[e[0]] = e[1]
	}
}

func getConfigValue(key, def string) string {
	if len(configs) == 0 {
		return def
	}
	if val, ok := configs[key]; ok {
		return val
	}
	return def
}

func MustGetConfigValue(key string) string {
	if len(configs) == 0 {
		panic("Config not parsed")
	}
	if val, ok := configs[key]; ok {
		return val
	}
	panic("Unknown config value " + key)
}
