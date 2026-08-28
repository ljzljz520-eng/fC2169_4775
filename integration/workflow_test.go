package integration

import (
	"campusawards/internal/domain"
	"campusawards/internal/ranking"
	"campusawards/internal/store"
	"campusawards/internal/vote"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	v := vote.NewService(s)
	c := domain.NewClub("c", "Top", "x", "")
	c.Status = "published"
	v.RegisterClub(c)
	v.RegisterVoter(domain.NewVoter("v", "student", "n", "d", true))
	if e := v.Submit(domain.NewVote("vote", "c", "v", "req")); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := domain.NewClub("c", "Top", "x", "")
	c.Status = "published"
	s.SaveClub(c)
	if _, e := ranking.NewService(s).Build(); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	if domain.StatusDescription("published") == "" {
		t.Fatal("missing")
	}
}
