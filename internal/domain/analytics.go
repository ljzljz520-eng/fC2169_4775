package domain

import "time"

type Period struct{ Start, End time.Time }

func NewPeriod(start, end time.Time) Period { return Period{Start: start, End: end} }
func (p Period) Contains(t time.Time) bool  { return !t.Before(p.Start) && !t.After(p.End) }
func ClubActive(c Club, p Period) bool      { return p.Contains(c.CreatedAt) && c.Status != "archived" }
func CountByStatus(clubs []Club) map[string]int {
	m := map[string]int{}
	for _, c := range clubs {
		m[NormalizeStatus(c.Status)]++
	}
	return m
}
func CountByCategory(clubs []Club) map[string]int {
	m := map[string]int{}
	for _, c := range clubs {
		m[c.Category]++
	}
	return m
}
func TotalClubVotes(clubs []Club) int {
	n := 0
	for _, c := range clubs {
		n += c.VoteCount
	}
	return n
}
func AverageVotes(clubs []Club) float64 {
	if len(clubs) == 0 {
		return 0
	}
	return float64(TotalClubVotes(clubs)) / float64(len(clubs))
}
func MaxVotes(clubs []Club) (Club, bool) {
	if len(clubs) == 0 {
		return Club{}, false
	}
	best := clubs[0]
	for _, c := range clubs[1:] {
		if c.VoteCount > best.VoteCount {
			best = c
		}
	}
	return best, true
}
func MinVotes(clubs []Club) (Club, bool) {
	if len(clubs) == 0 {
		return Club{}, false
	}
	best := clubs[0]
	for _, c := range clubs[1:] {
		if c.VoteCount < best.VoteCount {
			best = c
		}
	}
	return best, true
}
func VoteShare(c Club, all []Club) float64 {
	total := TotalClubVotes(all)
	if total == 0 {
		return 0
	}
	return float64(c.VoteCount) / float64(total)
}
func EligibleCount(voters []Voter) int {
	n := 0
	for _, v := range voters {
		if EligibleForVote(v) {
			n++
		}
	}
	return n
}
func ReviewOutcome(reviews []Review) map[string]int {
	m := map[string]int{}
	for _, r := range reviews {
		m[r.Decision]++
	}
	return m
}
func LatestReview(reviews []Review, club string) (Review, bool) {
	var out Review
	ok := false
	for _, r := range reviews {
		if r.ClubID == club && (!ok || r.CreatedAt.After(out.CreatedAt)) {
			out = r
			ok = true
		}
	}
	return out, ok
}
