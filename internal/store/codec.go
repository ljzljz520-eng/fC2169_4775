package store

import (
	"campusawards/internal/domain"
	"encoding/json"
)

func EncodeClub(c domain.Club) ([]byte, error) { return json.Marshal(c) }
func DecodeClub(b []byte) (domain.Club, error) {
	var c domain.Club
	e := json.Unmarshal(b, &c)
	return c, e
}
func EncodeVoter(v domain.Voter) ([]byte, error) { return json.Marshal(v) }
func DecodeVoter(b []byte) (domain.Voter, error) {
	var v domain.Voter
	e := json.Unmarshal(b, &v)
	return v, e
}
func EncodeVote(v domain.Vote) ([]byte, error) { return json.Marshal(v) }
func DecodeVote(b []byte) (domain.Vote, error) {
	var v domain.Vote
	e := json.Unmarshal(b, &v)
	return v, e
}
func EncodeReview(r domain.Review) ([]byte, error) { return json.Marshal(r) }
func DecodeReview(b []byte) (domain.Review, error) {
	var r domain.Review
	e := json.Unmarshal(b, &r)
	return r, e
}
func (s *Store) SaveManyClubs(clubs []domain.Club) error {
	for _, c := range clubs {
		if e := s.SaveClub(c); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) SaveManyVoters(voters []domain.Voter) error {
	for _, v := range voters {
		if e := s.SaveVoter(v); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) SaveManyVotes(votes []domain.Vote) error {
	for _, v := range votes {
		if e := s.SaveVote(v); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) SaveManyReviews(reviews []domain.Review) error {
	for _, r := range reviews {
		if e := s.SaveReview(r); e != nil {
			return e
		}
	}
	return nil
}
