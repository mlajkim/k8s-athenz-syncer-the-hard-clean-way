package syncer

import (
	"context"
	"fmt"
)

// TODO: Use Solution Template in the future, if possible.
func (s *Syncer) NsIntoAthenzDomain(ctx context.Context, ns string) error {
	// 1. CREATE SUB DOMAIN:
	newDomain := fmt.Sprintf("%s.%s", s.c.Syncer.ParentDomain, ns)
	if _, err := s.athenzClient.PostSubDomain(newDomain); err != nil {
		return fmt.Errorf("create subdomain failed: %w", err)
	}

	// 2. CREATE NECESSARY ROLES:
	defaultRoles := []string{s.c.Hard.AdminRoleName, s.c.Hard.ReadonlyRoleName}
	for _, role := range defaultRoles {
		if err := s.athenzClient.PostRole(newDomain, role); err != nil {
			return fmt.Errorf("create role %s failed: %w", role, err)
		}
	}

	// 3. CREATE BASIC ROLE in the namespace!
	

	return nil
}
