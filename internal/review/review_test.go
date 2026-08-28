package review

import (
	"campusawards/internal/domain"
	"campusawards/internal/store"
	"path/filepath"
	"testing"
)

func TestWorkflowReview(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := domain.NewClub("c", "社团", "艺体", "")
	s.SaveClub(c)
	if e := NewService(s).Decide(domain.NewReview("r", "c", "judge", "approved", "ok")); e != nil {
		t.Fatal(e)
	}
	got, _ := s.GetClub("c")
	if got.Status != "published" {
		t.Fatal(got.Status)
	}
}
