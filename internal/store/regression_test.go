package store

import (
	"fmt"
	"path/filepath"
	"sync"
	"testing"

	"campusawards/internal/domain"
	"go.etcd.io/bbolt"
)

// incrementOld reproduces the pre-fix read-modify-write race: read under a
// separate View tx, increment in memory, write under a separate Update tx, all
// guarded by RLock (shared). Two goroutines that race here lose a vote.
func (s *Store) incrementOld(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var c domain.Club
	if e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, "clubs", id, &c) }); e != nil {
		return fmt.Errorf("load club: %w", e)
	}
	c.VoteCount++
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "clubs", id, c) })
}

func TestIncrementOldLosesVotes(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "t.db"))
	if e != nil {
		t.Fatalf("open: %v", e)
	}
	defer s.Close()

	club := domain.NewClub("c1", "Robotics", "tech", "demo")
	club.Status = "published"
	if e := s.SaveClub(club); e != nil {
		t.Fatalf("save club: %v", e)
	}

	const voters = 50
	var wg sync.WaitGroup
	wg.Add(voters)
	for i := 0; i < voters; i++ {
		go func() {
			defer wg.Done()
			if e := s.incrementOld("c1"); e != nil {
				t.Errorf("increment: %v", e)
			}
		}()
	}
	wg.Wait()

	got, e := s.GetClub("c1")
	if e != nil {
		t.Fatalf("get club: %v", e)
	}
	// The old racy implementation drops votes; we expect fewer than 50.
	if got.VoteCount == voters {
		t.Fatalf("old implementation unexpectedly counted all %d votes; race did not trigger", voters)
	}
	t.Logf("old implementation counted %d of %d votes (lost %d) — confirms the bug", got.VoteCount, voters, voters-got.VoteCount)
}
