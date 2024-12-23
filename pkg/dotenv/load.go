package dotenv

import (
	"io/fs"
	"os"
	"strings"
)

func LoadEnv() {
	file, err := fs.ReadFile(os.DirFS("."), ".env")
	if err != nil {
		return
	}

	for _, line := range strings.Split(string(file), "\n") {
		parts := strings.Split(line, "=")

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}
}

func GetMissingVars(vars ...string) []string {
	var leftVars []string
	for _, v := range vars {
		if os.Getenv(v) == "" {
			return append(leftVars, v)
		}
	}
	return leftVars
}

func LoadDefaultIfNotEnv(key string, value string) {
	if _, ok := os.LookupEnv(key); !ok {
		os.Setenv(key, value)
	}
}
