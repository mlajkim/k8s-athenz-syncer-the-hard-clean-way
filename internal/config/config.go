package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Hard struct {
	AdminRoleName    string
	ReadonlyRoleName string
}

type Config struct {
	Athenz Athenz `yaml:"athenz"`
	Syncer Syncer `yaml:"syncer"`
	Hard   Hard
}

type Athenz struct {
	ZmsURL   string `yaml:"zmsUrl"`
	CertPath string `yaml:"certPath"`
	KeyPath  string `yaml:"keyPath"`
}

type Syncer struct {
	// Shared:
	ParentDomain string `yaml:"parentDomain"`

	// Specific:
	ARoleMembers ARoleMembers `yaml:"athenzRoleMembers"`
	Namespaces   Namespaces   `yaml:"namespaces"`
}

type ARoleMembers struct {
	Interval string `yaml:"interval"`
}

type Namespaces struct {
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

	cfg.hardCode() // For quick test & use the config as SSOT. nothing wrong.

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

	if c.Syncer.ParentDomain == "" {
		return fmt.Errorf("syncer.syncParentDomain is missing")
	}
	if c.Syncer.ARoleMembers.Interval == "" {
		return fmt.Errorf("syncer.athenzRoleMembers.interval is missing")
	}

	return nil
}

func (c *Config) hardCode() {
	c.Hard.AdminRoleName = "k8s_ns_admins"
	c.Hard.ReadonlyRoleName = "k8s_ns_viewers"
}
