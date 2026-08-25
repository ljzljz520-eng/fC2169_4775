package vote

import (
	"campusawards/internal/domain"
	"campusawards/internal/store"
	"path/filepath"
	"sync"
	"testing"
)

func TestConcurrentVotesCounted(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	svc := NewService(s)
	c := domain.NewClub("c", "社团", "科技", "")
	c.Status = "published"
	svc.RegisterClub(c)
	for i := 0; i < 2; i++ {
		svc.RegisterVoter(domain.NewVoter(string(rune('a'+i)), string(rune('s'+i)), "n", "d", true))
	}
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e := svc.Submit(domain.NewVote(string(rune('v'+i)), "c", string(rune('a'+i)), string(rune('r'+i))))
			if e != nil {
				t.Error(e)
			}
		}(i)
	}
	wg.Wait()
	got, _ := s.GetClub("c")
	if got.VoteCount != 2 {
		t.Fatalf("want 2 got %d", got.VoteCount)
	}
}
