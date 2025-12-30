package config

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
	ExcludedNamespaces    []string            `yaml:"excludedNamespaces"` // Raw
	ExcludedNamespacesMap map[string]struct{} `yaml:"-"`                  // Processed
	ARoleMembers          ARoleMembers        `yaml:"athenzRoleMembers"`
	Namespace             Namespaces          `yaml:"namespace"`
}

type RoleConfig struct {
	Suffix string       `yaml:"suffix"` // Role name is built based on given customizable suffix
	Rules  []PolicyRule `yaml:"rules"`
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
