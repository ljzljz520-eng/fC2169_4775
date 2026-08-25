package report

import "campusawards/internal/domain"

func TotalVotes(entries []domain.RankingEntry) int {
	n := 0
	for _, e := range entries {
		n += e.Votes
	}
	return n
}
func Leaders(entries []domain.RankingEntry, n int) []domain.RankingEntry {
	if n < 1 {
		return nil
	}
	if n > len(entries) {
		n = len(entries)
	}
	return entries[:n]
}
