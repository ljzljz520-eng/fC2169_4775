package store

import (
	"campusawards/internal/domain"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	c := domain.NewClub("c1", "星火", "科技", "")
	c.Status = "published"
	if e = s.SaveClub(c); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetClub("c1")
	if e != nil || got.Name != "星火" {
		t.Fatalf("%+v %v", got, e)
	}
}
