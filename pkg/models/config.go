package models

// Config is the global app config stored in ~/.ranpo/config.yaml.
type Config struct {
	ActiveEnv   string `yaml:"active_env"`
	DefaultAuth struct {
		Type  string `yaml:"type"`
		Token string `yaml:"token"`
	} `yaml:"default_auth"`
}
