package config

import (
	"fmt"
	"os"

	"github.com/mlajkim/k8s-athenz-syncer-the-hard-clean-way/pkg/util"
	"gopkg.in/yaml.v3"
)

// Load does:
// 1. Reads the config file from the given path
// 2. Unmarshals the content into Config struct
// 3. Builds derived state with the unmarshaled config
// 4. Finally Validates the config, including both original and derived state for correctness
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := cfg.buildDerivedState(nil); err != nil {
		return nil, fmt.Errorf("failed to build derived state: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) buildDerivedState(value *yaml.Node) error {
	c.Syncer.ExcludedNamespacesMap = util.StrArrayIntoUniqSet(c.Syncer.ExcludedNamespaces)

	return nil
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

	// TODO: Temporary testing:
	if c.Syncer.ExcludedNamespaces == nil || len(c.Syncer.ExcludedNamespaces) == 0 {
		return fmt.Errorf("syncer.excludedNamespaces is missing")
	}

	return nil
}
