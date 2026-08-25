package store

import (
	"campusawards/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"sync"
	"time"
)

var buckets = []string{"clubs", "voters", "votes", "reviews", "meta"}

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, &bbolt.Options{Timeout: time.Second})
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(b)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func (s *Store) Path() string { return s.path }
func put(tx *bbolt.Tx, bucket, key string, v any) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket([]byte(bucket)).Put([]byte(key), raw)
}
func get(tx *bbolt.Tx, bucket, key string, v any) error {
	raw := tx.Bucket([]byte(bucket)).Get([]byte(key))
	if raw == nil {
		return errors.New("not found")
	}
	return json.Unmarshal(raw, v)
}
func (s *Store) SaveClub(c domain.Club) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "clubs", c.ID, c) })
}
func (s *Store) GetClub(id string) (domain.Club, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var c domain.Club
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, "clubs", id, &c) })
	return c, e
}
func (s *Store) SaveVoter(v domain.Voter) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "voters", v.ID, v) })
}
func (s *Store) GetVoter(id string) (domain.Voter, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var v domain.Voter
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, "voters", id, &v) })
	return v, e
}
func (s *Store) SaveVote(v domain.Vote) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "votes", v.ID, v) })
}
func (s *Store) SaveReview(r domain.Review) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "reviews", r.ID, r) })
}
func (s *Store) ListClubs() ([]domain.Club, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []domain.Club{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("clubs")).ForEach(func(_, v []byte) error {
			var c domain.Club
			if e := json.Unmarshal(v, &c); e != nil {
				return e
			}
			out = append(out, c)
			return nil
		})
	})
	return out, e
}
func (s *Store) ListReviews() ([]domain.Review, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []domain.Review{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("reviews")).ForEach(func(_, v []byte) error {
			var r domain.Review
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) IncrementClubVotes(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	// Read, increment and write inside a single Update transaction. bbolt
	// serializes writers, so two concurrent votes for the same club can no
	// longer read the same stale count and overwrite each other, which used
	// to lose a vote when two voters voted the same club at once.
	return s.db.Update(func(tx *bbolt.Tx) error {
		var c domain.Club
		if e := get(tx, "clubs", id, &c); e != nil {
			return fmt.Errorf("load club: %w", e)
		}
		c.VoteCount++
		return put(tx, "clubs", id, c)
	})
}
