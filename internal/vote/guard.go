package vote

import (
	"campusawards/internal/domain"
	"errors"
)

var ErrRequestID = errors.New("request id required")

func CheckRequest(v domain.Vote) error {
	if v.RequestID == "" {
		return ErrRequestID
	}
	return nil
}
func (s *Service) ValidateSubmission(v domain.Vote) error {
	if e := CheckRequest(v); e != nil {
		return e
	}
	return s.EnsureReady()
}
func SameClub(a, b domain.Vote) bool { return a.ClubID == b.ClubID }
