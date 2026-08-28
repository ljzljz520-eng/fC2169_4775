package review

import (
	"campusawards/internal/domain"
	"campusawards/internal/store"
	"errors"
)

type Service struct{ Store *store.Store }

func NewService(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Decide(r domain.Review) error {
	if e := domain.ValidateReview(r); e != nil {
		return e
	}
	c, e := s.Store.GetClub(r.ClubID)
	if e != nil {
		return e
	}
	if !domain.CanReview(r, c) {
		return errors.New("review unavailable")
	}
	domain.ApplyReview(&c, r)
	if e = s.Store.SaveReview(r); e != nil {
		return e
	}
	return s.Store.SaveClub(c)
}
func (s *Service) Publish(id string) error {
	c, e := s.Store.GetClub(id)
	if e != nil {
		return e
	}
	c.Status = "published"
	return s.Store.SaveClub(c)
}
func (s *Service) Reject(id string) error {
	c, e := s.Store.GetClub(id)
	if e != nil {
		return e
	}
	c.Status = "rejected"
	return s.Store.SaveClub(c)
}
func (s *Service) History() ([]domain.Review, error) { return s.Store.ListReviews() }
func (s *Service) Approved(id string) bool {
	c, e := s.Store.GetClub(id)
	return e == nil && c.Status == "published"
}
