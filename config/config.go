package config

const (
	OPENROUTER_API_URL = "https://openrouter.ai/api/v1"
	DEFAULT_MODEL      = "x-ai/grok-4-fast:free"
)

type Config struct {
	API_URL       string
	API_KEY       string
	DefaultModel  string
	DefaultStream bool
}
