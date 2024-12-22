package dotenv

import "os"

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key); if !exists {
		value = defaultValue
	}
	return value
}