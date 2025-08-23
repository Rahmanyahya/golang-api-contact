package helper

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseEnvList(key string) []string {
	val, exist := os.LookupEnv(key)

	if !exist || val == "" {
		return []string{}
	}

	parts := strings.Split(val, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return parts
}

func GetEnvBool(key string, defaultValue bool) bool {
	val, exists := os.LookupEnv(key)

	if !exists || val == "" {
		return defaultValue
	}

	parsedVal, err := strconv.ParseBool(strings.ToLower(val))

	if err != nil {
		log.Printf("Warning: Couldn;t parse boolean value '%s' : %v. System using default:  %v", key, err, defaultValue)
	}

	return parsedVal
}