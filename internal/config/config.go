package config

import (
	"os"
	"path/filepath"

	"github.com/awiipp/ranpo/pkg/models"
	"gopkg.in/yaml.v3"
)

// Dir returns the ~/.ranpo directory path.
func Dir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".ranpo")
}

// Load reads config.yaml, returning defaults if the file doesn't exist.
func Load() (*models.Config, error) {
	path := filepath.Join(Dir(), "config.yaml")

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &models.Config{ActiveEnv: "local"}, nil
	}

	if err != nil {
		return nil, err
	}

	var cfg models.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Save writes config to disk, creating the directory if needed.
func Save(cfg *models.Config) error {
	if err := os.MkdirAll(Dir(), 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	
	return os.WriteFile(filepath.Join(Dir(), "config.yaml"), data, 0644)
}