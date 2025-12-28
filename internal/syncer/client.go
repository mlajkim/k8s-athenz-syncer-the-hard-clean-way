package syncer

import (
	"github.com/mlajkim/k8s-athenz-syncer-the-hard-clean-way/internal/config"
	"github.com/mlajkim/k8s-athenz-syncer-the-hard-clean-way/pkg/athenz"
)

// Syncer: Athenz와 K8s 사이의 로직을 담당
type Syncer struct {
	athenzClient *athenz.AthenzClient
	c            *config.Config
}

func New(client *athenz.AthenzClient, cfg *config.Config) *Syncer {
	return &Syncer{
		athenzClient: client,
		c:            cfg,
	}
}
