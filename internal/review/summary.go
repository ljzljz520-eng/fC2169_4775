package review

import (
	"campusawards/internal/domain"
	"sort"
)

type Summary struct{ Total, Approved, Rejected int }

func BuildSummary(reviews []domain.Review) Summary {
	out := Summary{}
	for _, r := range reviews {
		out.Total++
		if r.Decision == "approved" {
			out.Approved++
		}
		if r.Decision == "rejected" {
			out.Rejected++
		}
	}
	return out
}
func ApprovalRate(s Summary) float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Approved) / float64(s.Total)
}
func ReviewerCounts(reviews []domain.Review) map[string]int {
	m := map[string]int{}
	for _, r := range reviews {
		m[r.ReviewerID]++
	}
	return m
}
func ClubsReviewed(reviews []domain.Review) []string {
	m := map[string]bool{}
	for _, r := range reviews {
		m[r.ClubID] = true
	}
	out := []string{}
	for id := range m {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}
func SortByTime(reviews []domain.Review) []domain.Review {
	out := append([]domain.Review(nil), reviews...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out
}
func Latest(reviews []domain.Review) domain.Review {
	if len(reviews) == 0 {
		return domain.Review{}
	}
	out := reviews[0]
	for _, r := range reviews[1:] {
		if r.CreatedAt.After(out.CreatedAt) {
			out = r
		}
	}
	return out
}
