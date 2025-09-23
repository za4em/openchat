package config

var (
	XAI_API_URL   = "https://api.x.ai/v1"
	DEFAULT_MODEL = "grok-4-fast-non-reasoning"
)

type Config struct {
	API_URL      string
	API_KEY      string
	DefaultModel string
}
