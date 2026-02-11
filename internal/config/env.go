package config

import "os"

func Environment() string {
	if env := os.Getenv("ENVIRONMENT"); env != "" {
		return env
	}
	return "production"
}

func IsDevelopment() bool {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	return env == "development"
}
