package ranking

import (
	"campusawards/internal/domain"
	"campusawards/internal/store"
)

type Service struct{ Store *store.Store }

func NewService(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Build() ([]domain.RankingEntry, error) {
	cs, e := s.Store.ListClubs()
	if e != nil {
		return nil, e
	}
	out := []domain.RankingEntry{}
	for _, c := range cs {
		if domain.IsPublished(c) {
			out = append(out, domain.RankingEntry{ClubID: c.ID, Name: c.Name, Category: c.Category, Votes: c.VoteCount})
		}
	}
	domain.SortRankings(out)
	return out, nil
}
func (s *Service) TopTen() ([]domain.RankingEntry, error) {
	r, e := s.Build()
	if e != nil {
		return nil, e
	}
	return domain.Top(r, 10), nil
}
func (s *Service) Category(name string) ([]domain.RankingEntry, error) {
	r, e := s.Build()
	if e != nil {
		return nil, e
	}
	out := r[:0]
	for _, x := range r {
		if x.Category == name {
			out = append(out, x)
		}
	}
	for i := range out {
		out[i].Position = i + 1
	}
	return out, nil
}
