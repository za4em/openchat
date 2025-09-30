package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/za4em/openchat/domain"
)

const (
	FILE_PERMISSION = 0644
)

type ChatStorage struct {
	configDir string
	Chats     map[string]domain.Chat
}

func NewChatStorage(configDir string) (*ChatStorage, error) {
	storage := &ChatStorage{
		configDir: configDir,
		Chats:     make(map[string]domain.Chat),
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
			storage.Chats[chat.ID] = *chat
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

func (storage *ChatStorage) Save(chat *domain.Chat) error {
	filePath := filepath.Join(storage.configDir, chat.ID+".json")
	data, err := json.MarshalIndent(chat, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, FILE_PERMISSION)
}
