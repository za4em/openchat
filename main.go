package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/za4em/openchat/config"
	"github.com/za4em/openchat/data/datasource/api"
	"github.com/za4em/openchat/data/datasource/storage"
	"github.com/za4em/openchat/data/store"
	"github.com/za4em/openchat/ui"
)

func main() {
	godotenv.Load(".env")
	apiKey := os.Getenv("API_KEY")
	config := config.Config{
		API_URL:       config.OPENROUTER_API_URL,
		API_KEY:       apiKey,
		DefaultModel:  config.DEFAULT_MODEL,
		DefaultStream: false,
	}
	api := api.NewOpenRouterApi(config)
	storage, error := storage.NewChatStorage()
	if error != nil {
		log.Fatal(error)
		return
	}
	store := &store.ChatStore{
		Api:     api,
		Storage: storage,
	}
	uiModel := ui.NewModel(store)
	p := tea.NewProgram(uiModel)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
