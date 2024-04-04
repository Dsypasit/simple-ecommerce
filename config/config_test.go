package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringEnv(t *testing.T) {
	// TODO: get string env
	tests := []struct {
		name       string
		key        string
		defaultVal string
		osVal      string
		want       string
	}{
		{
			name:       "get str from env",
			key:        "host",
			defaultVal: "localhost",
			osVal:      "127.0.0.1",
			want:       "127.0.0.1",
		},
		{
			name:       "get default value when empty env",
			key:        "host",
			defaultVal: "localhost",
			osVal:      "",
			want:       "localhost",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Setenv(test.key, test.osVal)

			value := envString(test.key, test.defaultVal)
			assert.Equal(t, test.want, value)
		})
	}
}

func TestIntEnv(t *testing.T) {
	// TODO: get int env
	tests := []struct {
		name       string
		key        string
		osVal      string
		defaultVal int
		want       int
	}{
		{
			name:       "get int from env",
			key:        "host",
			defaultVal: 8000,
			osVal:      "2000",
			want:       2000,
		},
		{
			name:       "get default value when empty env",
			key:        "host",
			defaultVal: 8000,
			osVal:      "",
			want:       8000,
		},
		{
			name:       "get default value when get strinv from env",
			key:        "host",
			defaultVal: 8000,
			osVal:      "wow",
			want:       8000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Setenv(test.key, test.osVal)

			value := envInt(test.key, test.defaultVal)
			assert.Equal(t, test.want, value)
		})
	}
}
