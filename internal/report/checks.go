package report

import (
	"campusawards/internal/domain"
	"fmt"
)

func CheckUnique(entries []domain.RankingEntry) error {
	seen := map[string]bool{}
	for _, e := range entries {
		if seen[e.ClubID] {
			return fmt.Errorf("duplicate club %s", e.ClubID)
		}
		seen[e.ClubID] = true
	}
	return nil
}
func CheckCategories(entries []domain.RankingEntry) error {
	for _, e := range entries {
		if e.Category == "" {
			return fmt.Errorf("missing category for %s", e.Name)
		}
	}
	return nil
}
func Quality(entries []domain.RankingEntry) error {
	if e := Validate(entries); e != nil {
		return e
	}
	if e := CheckUnique(entries); e != nil {
		return e
	}
	return CheckCategories(entries)
}
