package env

import (
	"testing"
	"os"
)

func TestString(t *testing.T) {
	t.Run("returns string env variable", func(t *testing.T) {
		envKey, envValue := "DB_HOST", "db.example.com"
		if err := os.Setenv(envKey, envValue); err != nil {
			t.Fatalf("failed to set env variable '%s': %s", envKey, err)
		}

		actual := String("DB_HOST", "localhost")
		if actual != envValue {
			t.Errorf("expected '%s' env variable to have valid '%s', got '%s'", envKey, envValue, actual)
		}
	})

	t.Run("returns default valid", func(t *testing.T) {
		envKey, defaultValue := "DB_HOST", "db.example.com"
		actual := String(envKey, defaultValue)
		if actual != defaultValue {
			t.Errorf("expected '%s' env variable to have valid '%s', got '%s'", envKey, defaultValue, actual)
		}
	})
}

func TestBool(t *testing.T) {
	t.Run("returns bool env variable", func(t *testing.T) {
		var tests = []struct {
			env   string
			valid bool
		}{
			{env: "TRUE", valid: true},
			{env: "true", valid: true},
			{env: "1", valid: true},
			{env: "F", valid: false},
			{env: "T", valid: false},
			{env: "False", valid: false},
			{env: "True", valid: false},
		}
		for _, testCase := range tests {
			t.Run(testCase.env, func(t *testing.T) {
				key := "DEBUG_MODE"
				if err := os.Setenv(key, testCase.env); err != nil {
					t.Fatalf("failed to set env var %s for test: %s\n", key, err)
				}

				def := false
				actual, expected := Bool(key, def), testCase.valid
				if  actual != expected {
					t.Errorf("expected %v, got %v", expected, actual)
				}
			})
		}
	})

	t.Run("returns default valid", func(t *testing.T) {
		def := true
		expected, actual := Bool("DEBUG_MODE", def), def
		if expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	})
}