package main

import (
	"log"

	"github.com/za4em/openchat/config"
	"github.com/za4em/openchat/datasource/api"
	"github.com/za4em/openchat/datasource/storage"
	"github.com/za4em/openchat/presentation"
	"github.com/za4em/openchat/store"
)

func main() {
	config := config.Config{
		API_URL:       config.XAI_API_URL,
		API_KEY:       "",
		DefaultModel:  config.DEFAULT_MODEL,
		DefaultStream: false,
	}
	api := api.CreateXaiApi(config)
	storage, error := storage.CreateChatStorage()
	if error != nil {
		log.Fatal(error)
		return
	}
	store := store.ChatStore{
		Api:     api,
		Storage: storage,
	}
	uiModel := presentation.UiModel{
		ChatStore: &store,
	}
	_ = uiModel
}
