package config

import "os"

func GetEnv(key, systemVal string) string {

	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}

	return systemVal

}