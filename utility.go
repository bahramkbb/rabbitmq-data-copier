package main

import "os"

// Get os env variable or return fallback
func GetEnvVariable(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

