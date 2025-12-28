package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Athenz AthenzConfig `yaml:"athenz"`
	Syncer SyncerConfig `yaml:"syncer"`
}

type AthenzConfig struct {
	ZmsURL   string `yaml:"zmsUrl"`
	CertPath string `yaml:"certPath"`
	KeyPath  string `yaml:"keyPath"`
}

type SyncerConfig struct {
	SyncParentDomain string           `yaml:"syncParentDomain"`
	AthenzRoleSyncer AthenzRoleSyncer `yaml:"athenzRoleSyncer"`
	// NamespaceSyncer is not yet required
}

type AthenzRoleSyncer struct {
	Interval string `yaml:"interval"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) validate() error {
	if c.Athenz.ZmsURL == "" {
		return fmt.Errorf("athenz.zmsUrl is missing")
	}
	if c.Athenz.CertPath == "" {
		return fmt.Errorf("athenz.certPath is missing")
	}
	if c.Athenz.KeyPath == "" {
		return fmt.Errorf("athenz.keyPath is missing")
	}

	if c.Syncer.SyncParentDomain == "" {
		return fmt.Errorf("syncer.syncParentDomain is missing")
	}
	if c.Syncer.AthenzRoleSyncer.Interval == "" {
		return fmt.Errorf("syncer.athenzRoleSyncer.interval is missing")
	}

	return nil
}
