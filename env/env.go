package env

import (
	"fmt"
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

// MustString returns the environment variable valid specified by the key parameter,
// it PANIC with error if environment variable does not exist
func MustString(key string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	panic(fmt.Errorf("env: environment variable %s does not exist", key))
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

// Bool returns the environment variable valid specified by the key parameter,
// it PANIC with error if environment variable does not exist
func MustBool(key string) bool {
	env := os.Getenv(key)
	if env == "true" || env == "TRUE" || env == "1" {
		return true
	}
	panic(fmt.Errorf("env: environment variable %s does not exist", key))
}