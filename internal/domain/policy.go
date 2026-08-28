package domain

import "sort"

func EligibleForVote(v Voter) bool { return v.Eligible && v.StudentID != "" }
func CanReview(r Review, c Club) bool {
	return r.ReviewerID != "" && c.ID == r.ClubID && c.Status != "archived"
}
func ApplyReview(c *Club, r Review) {
	if r.Decision == "approved" {
		c.Status = "published"
	} else {
		c.Status = "rejected"
	}
}
func SortRankings(entries []RankingEntry) {
	sort.SliceStable(entries, func(i, j int) bool {
		if entries[i].Votes == entries[j].Votes {
			return entries[i].Name < entries[j].Name
		}
		return entries[i].Votes > entries[j].Votes
	})
	for i := range entries {
		entries[i].Position = i + 1
	}
}
func Top(entries []RankingEntry, n int) []RankingEntry {
	if n < 1 {
		return nil
	}
	if n > len(entries) {
		n = len(entries)
	}
	return append([]RankingEntry(nil), entries[:n]...)
}
func CategoryAllowed(c Club, allowed map[string]bool) bool {
	if len(allowed) == 0 {
		return true
	}
	return allowed[c.Category]
}
