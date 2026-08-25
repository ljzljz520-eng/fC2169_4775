package report

import (
	"campusawards/internal/domain"
	"sort"
)

func ByCategory(entries []domain.RankingEntry) map[string][]domain.RankingEntry {
	m := map[string][]domain.RankingEntry{}
	for _, e := range entries {
		m[e.Category] = append(m[e.Category], e)
	}
	for k := range m {
		sort.Slice(m[k], func(i, j int) bool { return m[k][i].Votes > m[k][j].Votes })
	}
	return m
}
func Winners(entries []domain.RankingEntry) []string {
	out := []string{}
	for _, e := range entries {
		if e.Position <= 10 {
			out = append(out, e.Name)
		}
	}
	return out
}
