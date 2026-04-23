package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/awiipp/ranpo/internal/config"
	"github.com/awiipp/ranpo/pkg/models"
)

func EnvDir() string {
	return filepath.Join(config.Dir(), "environments")
}

func SaveEnv(env *models.Environtment) error {
	if err := os.MkdirAll(EnvDir(), 0755); err != nil {
		return nil
	}

	data, err := json.MarshalIndent(env, "", "  ")
	if err != nil {
		return nil
	}

	return os.WriteFile(filepath.Join(EnvDir(), sanitize(env.Name)+".json"), data, 0644)
}

func LoadEnv(name string) (*models.Environtment, error) {
	path := filepath.Join(EnvDir(), sanitize(name)+".json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var env models.Environtment
	return &env, json.Unmarshal(data, &env)
}

func ListEnvs() ([]string, error) {
	entries, err := os.ReadDir(EnvDir())
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
