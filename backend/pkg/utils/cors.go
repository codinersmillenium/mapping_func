package utils

import (
	"mapping_func/config"
	"strings"
)

func CorsOrigins(cfg *config.Config) []string {
	if cfg.CorsAllowedOrigins == "*" {
		return []string{"*"}
	}
	return strings.Split(cfg.CorsAllowedOrigins, ",")
}
