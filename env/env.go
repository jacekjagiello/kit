package env

import (
	"os"
)

// String returns the environment variable valid specified by the key parameter,
// otherwise returning a default valid if set.
func String(key, def string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	return def
}

// Bool returns the environment variable valid specified by the key parameter,
// otherwise returning a default valid if set.
func Bool(key string, def bool) bool {
	env := os.Getenv(key)
	if env == "true" || env == "TRUE" || env == "1" {
		return true
	}
	return def
}
