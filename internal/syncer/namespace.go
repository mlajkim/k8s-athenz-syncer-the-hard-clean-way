package syncer

import (
	"context"
	"fmt"
)

// TODO: Use Solution Template in the future, if possible.
func (s *Syncer) NsIntoAthenzDomain(ctx context.Context, ns string) error {
	newDomain := fmt.Sprintf("%s.%s", s.c.Syncer.ParentDomain, ns)
	if _, err := s.athenzClient.PostSubDomain(newDomain); err != nil {
		return fmt.Errorf("create subdomain failed: %w", err)
	}

	// Create necessary roles:
	defaultRoles := []string{"k8s_ns_admins", "k8s_ns_viewers"}
	for _, role := range defaultRoles {
		if err := s.athenzClient.PostRole(newDomain, role); err != nil {
			return fmt.Errorf("create role %s failed: %w", role, err)
		}
	}

	return nil
}
