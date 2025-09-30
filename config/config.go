package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

const (
	OPENROUTER_API_URL = "https://openrouter.ai/api/v1"
	DEFAULT_MODEL      = "x-ai/grok-4-fast:free"
	APP_PATH           = "openchat"
	CHAT_PATH          = "chats"
	DIR_PERMISSION     = 0755
)

type Config struct {
	API_URL       string
	API_KEY       string
	DefaultModel  string
	DefaultStream bool
	ConfigDir     string
}

func CreateConfigDir() (string, error) {
	configDir := filepath.Join(xdg.ConfigHome, APP_PATH, CHAT_PATH)
	err := os.MkdirAll(configDir, os.FileMode(DIR_PERMISSION))
	return configDir, err
}
