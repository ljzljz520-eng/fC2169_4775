package report

import (
	"campusawards/internal/domain"
	"fmt"
	"strings"
)

func Export(entries []domain.RankingEntry, title string) string {
	var b strings.Builder
	fmt.Fprintln(&b, title)
	for _, e := range entries {
		fmt.Fprintf(&b, "%02d %s %s %d票\n", e.Position, e.Name, e.Category, e.Votes)
	}
	return b.String()
}
func Summary(entries []domain.RankingEntry) map[string]int {
	m := map[string]int{}
	for _, e := range entries {
		m[e.Category]++
	}
	return m
}
func Validate(entries []domain.RankingEntry) error {
	if len(entries) == 0 {
		return fmt.Errorf("empty ranking")
	}
	for i, e := range entries {
		if e.Position != i+1 {
			return fmt.Errorf("position gap")
		}
	}
	return nil
}
