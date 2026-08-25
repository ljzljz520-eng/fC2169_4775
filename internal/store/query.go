package store

import (
	"campusawards/internal/domain"
	"strings"
)

func (s *Store) FindClubs(term string) ([]domain.Club, error) {
	cs, e := s.ListClubs()
	if e != nil {
		return nil, e
	}
	term = strings.ToLower(strings.TrimSpace(term))
	out := []domain.Club{}
	for _, c := range cs {
		if term == "" || strings.Contains(strings.ToLower(c.Name), term) || strings.Contains(strings.ToLower(c.Category), term) {
			out = append(out, c)
		}
	}
	return out, nil
}
func (s *Store) CountPublished() int {
	cs, e := s.ListClubs()
	if e != nil {
		return 0
	}
	n := 0
	for _, c := range cs {
		if c.Status == "published" {
			n++
		}
	}
	return n
}
func (s *Store) ExistsClub(id string) bool { _, e := s.GetClub(id); return e == nil }
