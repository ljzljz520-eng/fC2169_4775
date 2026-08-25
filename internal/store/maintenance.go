package store

import (
	"campusawards/internal/domain"
	"encoding/json"
	"go.etcd.io/bbolt"
	"time"
)

type Snapshot struct {
	Clubs   []domain.Club
	Reviews []domain.Review
	TakenAt time.Time
}

func (s *Store) Snapshot() (Snapshot, error) {
	cs, e := s.ListClubs()
	if e != nil {
		return Snapshot{}, e
	}
	rs, e := s.ListReviews()
	if e != nil {
		return Snapshot{}, e
	}
	return Snapshot{Clubs: cs, Reviews: rs, TakenAt: time.Now().UTC()}, nil
}
func (s *Store) ExportJSON() ([]byte, error) {
	snap, e := s.Snapshot()
	if e != nil {
		return nil, e
	}
	return json.MarshalIndent(snap, "", "  ")
}
func (s *Store) ImportClubs(clubs []domain.Club) error {
	for _, c := range clubs {
		if e := s.SaveClub(c); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) PublishedNames() []string {
	cs, e := s.ListClubs()
	if e != nil {
		return nil
	}
	out := []string{}
	for _, c := range cs {
		if c.Status == "published" {
			out = append(out, c.Name)
		}
	}
	return out
}
func (s *Store) ResetVotes() error {
	cs, e := s.ListClubs()
	if e != nil {
		return e
	}
	for _, c := range cs {
		c.VoteCount = 0
		if e = s.SaveClub(c); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) TouchMeta(key string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("meta")).Put([]byte(key), []byte(time.Now().UTC().Format(time.RFC3339)))
	})
}
