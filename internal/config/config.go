package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Athenz Athenz `yaml:"athenz"`
	Syncer Syncer `yaml:"syncer"`
}

type Athenz struct {
	ZmsURL   string `yaml:"zmsUrl"`
	CertPath string `yaml:"certPath"`
	KeyPath  string `yaml:"keyPath"`
}

type Syncer struct {
	// Shared:
	ParentDomain string       `yaml:"parentDomain"`
	Roles        []RoleConfig `yaml:"roles"`

	// Specific:
	ARoleMembers ARoleMembers `yaml:"athenzRoleMembers"`
	Namespace    Namespaces   `yaml:"namespace"`
}

type RoleConfig struct {
	Name  string       `yaml:"name"` // Both ATHENZ ROLE NAME & K8S ROLE NAME
	Rules []PolicyRule `yaml:"rules"`
}

type PolicyRule struct {
	APIGroups []string `yaml:"apiGroups"`
	Resources []string `yaml:"resources"`
	Verbs     []string `yaml:"verbs"`
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
