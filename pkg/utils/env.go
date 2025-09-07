package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvAsInt(key string, defaultValue int) (int, error) {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue, nil
	}
	i, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("invalid value for %s: %v", key, err)
	}
	return i, nil
}

func GetEnvAsBool(key string, defaultValue bool) (bool, error) {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue, nil
	}
	b, err := strconv.ParseBool(valueStr)
	if err != nil {
		return false, fmt.Errorf("invalid value for %s: %v", key, err)
	}
	return b, nil
}

func GetEnvAsSlice(key string, defaultValue []string, sep string) []string {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	s := strings.Split(valueStr, sep)
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}
	return s
}
