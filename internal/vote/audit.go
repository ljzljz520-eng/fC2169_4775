package vote

import (
	"campusawards/internal/domain"
	"campusawards/internal/store"
	"sort"
)

type Audit struct{ Store *store.Store }

func NewAudit(s *store.Store) *Audit              { return &Audit{Store: s} }
func (a *Audit) ClubVotes(id string) (int, error) { c, e := a.Store.GetClub(id); return c.VoteCount, e }
func (a *Audit) PublishedClubs() ([]domain.Club, error) {
	cs, e := a.Store.ListClubs()
	if e != nil {
		return nil, e
	}
	out := cs[:0]
	for _, c := range cs {
		if domain.IsPublished(c) {
			out = append(out, c)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}
func (a *Audit) VoteSummary() map[string]int {
	m := map[string]int{}
	if cs, e := a.Store.ListClubs(); e == nil {
		for _, c := range cs {
			m[c.ID] = c.VoteCount
		}
	}
	return m
}
