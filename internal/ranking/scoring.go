package ranking

import "campusawards/internal/domain"

type ScoreCard struct {
	ClubID                                   string
	Votes, ReviewScore, ActivityScore, Total int
}

func NewScoreCard(c domain.Club) ScoreCard { return ScoreCard{ClubID: c.ID, Votes: c.VoteCount} }
func (s *ScoreCard) AddReview(points int) {
	if points < 0 {
		points = 0
	}
	if points > 100 {
		points = 100
	}
	s.ReviewScore = points
	s.recompute()
}
func (s *ScoreCard) AddActivity(points int) {
	if points < 0 {
		points = 0
	}
	if points > 100 {
		points = 100
	}
	s.ActivityScore = points
	s.recompute()
}
func (s *ScoreCard) AddVotes(v int) {
	if v > 0 {
		s.Votes = v
	}
	s.recompute()
}
func (s *ScoreCard) recompute()    { s.Total = s.Votes + s.ReviewScore + s.ActivityScore }
func (s ScoreCard) Eligible() bool { return s.ReviewScore >= 60 && s.ActivityScore >= 50 }
func (s ScoreCard) Tier() string {
	switch {
	case s.Total >= 200:
		return "gold"
	case s.Total >= 150:
		return "silver"
	case s.Total >= 100:
		return "bronze"
	default:
		return "candidate"
	}
}
func BuildScoreCards(clubs []domain.Club) []ScoreCard {
	out := make([]ScoreCard, 0, len(clubs))
	for _, c := range clubs {
		out = append(out, NewScoreCard(c))
	}
	return out
}
func ScoreByVotes(cards []ScoreCard) []ScoreCard {
	out := append([]ScoreCard(nil), cards...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Votes > out[i].Votes {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
func ScoreByTotal(cards []ScoreCard) []ScoreCard {
	out := append([]ScoreCard(nil), cards...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Total > out[i].Total {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
func NormalizeCards(cards []ScoreCard) []ScoreCard {
	out := append([]ScoreCard(nil), cards...)
	for i := range out {
		if out[i].Votes < 0 {
			out[i].Votes = 0
		}
		out[i].recompute()
	}
	return out
}
func TopCards(cards []ScoreCard, n int) []ScoreCard {
	if n < 0 {
		return nil
	}
	if n > len(cards) {
		n = len(cards)
	}
	return append([]ScoreCard(nil), ScoreByTotal(cards)[:n]...)
}
func AverageScore(cards []ScoreCard) float64 {
	if len(cards) == 0 {
		return 0
	}
	sum := 0
	for _, c := range cards {
		sum += c.Total
	}
	return float64(sum) / float64(len(cards))
}
func HighestScore(cards []ScoreCard) (ScoreCard, bool) {
	if len(cards) == 0 {
		return ScoreCard{}, false
	}
	best := cards[0]
	for _, c := range cards[1:] {
		if c.Total > best.Total {
			best = c
		}
	}
	return best, true
}
func LowestScore(cards []ScoreCard) (ScoreCard, bool) {
	if len(cards) == 0 {
		return ScoreCard{}, false
	}
	best := cards[0]
	for _, c := range cards[1:] {
		if c.Total < best.Total {
			best = c
		}
	}
	return best, true
}
func CountEligible(cards []ScoreCard) int {
	n := 0
	for _, c := range cards {
		if c.Eligible() {
			n++
		}
	}
	return n
}
func CountTier(cards []ScoreCard) map[string]int {
	m := map[string]int{}
	for _, c := range cards {
		m[c.Tier()]++
	}
	return m
}
func FilterTier(cards []ScoreCard, tier string) []ScoreCard {
	out := []ScoreCard{}
	for _, c := range cards {
		if c.Tier() == tier {
			out = append(out, c)
		}
	}
	return out
}
func VoteContribution(c ScoreCard) float64 {
	if c.Total == 0 {
		return 0
	}
	return float64(c.Votes) / float64(c.Total)
}
func ReviewContribution(c ScoreCard) float64 {
	if c.Total == 0 {
		return 0
	}
	return float64(c.ReviewScore) / float64(c.Total)
}
func ActivityContribution(c ScoreCard) float64 {
	if c.Total == 0 {
		return 0
	}
	return float64(c.ActivityScore) / float64(c.Total)
}
func Compare(a, b ScoreCard) int {
	if a.Total > b.Total {
		return 1
	}
	if a.Total < b.Total {
		return -1
	}
	if a.Votes > b.Votes {
		return 1
	}
	if a.Votes < b.Votes {
		return -1
	}
	return 0
}
func RankCards(cards []ScoreCard) []ScoreCard {
	out := ScoreByTotal(cards)
	for i := range out {
		out[i].Total = out[i].Total
	}
	return out
}
func AddBonus(cards []ScoreCard, bonus map[string]int) []ScoreCard {
	out := append([]ScoreCard(nil), cards...)
	for i := range out {
		if b, ok := bonus[out[i].ClubID]; ok {
			out[i].ActivityScore += b
			out[i].recompute()
		}
	}
	return out
}
func CapScores(cards []ScoreCard, max int) []ScoreCard {
	out := append([]ScoreCard(nil), cards...)
	for i := range out {
		if out[i].Total > max {
			out[i].Total = max
		}
	}
	return out
}
func CopyCards(cards []ScoreCard) []ScoreCard { return append([]ScoreCard(nil), cards...) }
func CardIDs(cards []ScoreCard) []string {
	ids := make([]string, 0, len(cards))
	for _, c := range cards {
		ids = append(ids, c.ClubID)
	}
	return ids
}
