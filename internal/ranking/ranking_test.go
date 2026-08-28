package ranking

import (
	"campusawards/internal/domain"
	"campusawards/internal/store"
	"path/filepath"
	"testing"
)

func TestWorkflowRanking(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	for i := 0; i < 2; i++ {
		c := domain.NewClub(string(rune('a'+i)), string(rune('A'+i)), "x", "")
		c.Status = "published"
		c.VoteCount = i
		s.SaveClub(c)
	}
	r, e := NewService(s).Build()
	if e != nil || len(r) != 2 || r[0].Votes != 1 {
		t.Fatalf("%v %+v", e, r)
	}
}
