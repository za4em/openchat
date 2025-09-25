package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/za4em/openchat/domain"
)

const (
	APP_PATH        = "openchat"
	CHAT_PATH       = "chats"
	DIR_PERMISSION  = 0755
	FILE_PERMISSION = 0644
)

type ChatStorage struct {
	configDir string
	Chats     map[string]*domain.Chat
}

func NewChatStorage() (*ChatStorage, error) {
	configDir, err := createConfigDir()
	if err != nil {
		return nil, err
	}

	storage := &ChatStorage{
		configDir: configDir,
		Chats:     make(map[string]*domain.Chat),
	}

	files, err := os.ReadDir(configDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			chat, err := storage.parseFile(file)
			if err != nil {
				continue
			}
			storage.Chats[chat.ID] = chat
		}
	}

	return storage, nil
}

func (storage *ChatStorage) parseFile(file os.DirEntry) (*domain.Chat, error) {
	filePath := filepath.Join(storage.configDir, file.Name())
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var chat domain.Chat
	err = json.Unmarshal(data, &chat)
	return &chat, err
}

func createConfigDir() (string, error) {
	configDir := filepath.Join(xdg.ConfigHome, APP_PATH, CHAT_PATH)
	err := os.MkdirAll(configDir, os.FileMode(DIR_PERMISSION))
	return configDir, err
}

func (storage *ChatStorage) Save(chat *domain.Chat) error {
	filePath := filepath.Join(storage.configDir, chat.ID+".json")
	data, err := json.MarshalIndent(chat, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, FILE_PERMISSION)
}
