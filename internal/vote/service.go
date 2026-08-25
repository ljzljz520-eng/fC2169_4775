package vote

import (
	"campusawards/internal/domain"
	"campusawards/internal/store"
	"errors"
	"fmt"
	"sync"
)

type Service struct {
	Store *store.Store
	mu    sync.Mutex
}

func NewService(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) RegisterVoter(v domain.Voter) error {
	if e := domain.ValidateVoter(v); e != nil {
		return e
	}
	return s.Store.SaveVoter(v)
}
func (s *Service) RegisterClub(c domain.Club) error {
	if e := domain.ValidateClub(c); e != nil {
		return e
	}
	return s.Store.SaveClub(c)
}
func (s *Service) Submit(v domain.Vote) error {
	if e := domain.ValidateVote(v); e != nil {
		return e
	}
	voter, e := s.Store.GetVoter(v.VoterID)
	if e != nil {
		return e
	}
	if !domain.EligibleForVote(voter) {
		return errors.New("voter not eligible")
	}
	club, e := s.Store.GetClub(v.ClubID)
	if e != nil {
		return e
	}
	if !domain.IsPublished(club) {
		return errors.New("club not published")
	}
	if e = s.Store.SaveVote(v); e != nil {
		return e
	}
	return s.Store.IncrementClubVotes(v.ClubID)
}
func (s *Service) SubmitSerialized(v domain.Vote) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Submit(v)
}
func (s *Service) Explain(v domain.Vote) string {
	if v.RequestID == "" {
		return "missing request id"
	}
	return fmt.Sprintf("request %s for club %s", v.RequestID, v.ClubID)
}
func (s *Service) EnsureReady() error {
	if s.Store == nil {
		return errors.New("store missing")
	}
	return nil
}
