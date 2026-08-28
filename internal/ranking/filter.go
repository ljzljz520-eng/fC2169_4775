package ranking

import "campusawards/internal/domain"

func Filter(entries []domain.RankingEntry, min int) []domain.RankingEntry {
	out := []domain.RankingEntry{}
	for _, e := range entries {
		if e.Votes >= min {
			out = append(out, e)
		}
	}
	for i := range out {
		out[i].Position = i + 1
	}
	return out
}
func Names(entries []domain.RankingEntry) []string {
	out := make([]string, 0, len(entries))
	for _, e := range entries {
		out = append(out, e.Name)
	}
	return out
}
