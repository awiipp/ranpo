package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/awiipp/ranpo/internal/config"
	"github.com/awiipp/ranpo/pkg/models"
)

func CollectionDir() string {
	return filepath.Join(config.Dir(), "collections")
}

func SaveCollection(c *models.Collection) error {
	if err := os.MkdirAll(CollectionDir(), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	name := sanitize(c.Name)
	return os.WriteFile(filepath.Join(CollectionDir(), name+".json"), data, 0644)
}

func LoadCollection(name string) (*models.Collection, error) {
	path := filepath.Join(CollectionDir(), sanitize(name)+".json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c models.Collection
	return &c, json.Unmarshal(data, &c)
}

func ListCollections() ([]string, error) {
	entries, err := os.ReadDir(CollectionDir())
	if err != nil {
		return nil, nil
	}
	
	var names []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".json") {
			names = append(names, strings.TrimSuffix(e.Name(), ".json"))
		}
	}
	
	return names, nil
}

func DeleteCollection(name string) error {
	return os.Remove(filepath.Join(CollectionDir(), sanitize(name)+".json"))
}