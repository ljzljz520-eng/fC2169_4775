package store

import (
	"path/filepath"
	"sync"
	"testing"

	"campusawards/internal/domain"
)

// TestIncrementClubVotesConcurrent verifies that two voters voting the same
// club at the same time each have their vote counted. Previously the increment
// read and wrote in two separate transactions under an RLock, so concurrent
// increments read the same stale count and clobbered each other, losing votes.
func TestIncrementClubVotesConcurrent(t *testing.T) {
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
			if e := s.IncrementClubVotes("c1"); e != nil {
				t.Errorf("increment: %v", e)
			}
		}()
	}
	wg.Wait()

	got, e := s.GetClub("c1")
	if e != nil {
		t.Fatalf("get club: %v", e)
	}
	if got.VoteCount != voters {
		t.Fatalf("vote count = %d, want %d (lost %d votes)", got.VoteCount, voters, voters-got.VoteCount)
	}
}
