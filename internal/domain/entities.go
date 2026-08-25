package domain

import (
	"errors"
	"strings"
	"time"
)

type Club struct {
	ID, Name, Category, Description, Status string
	VoteCount                               int
	CreatedAt                               time.Time
}
type Voter struct {
	ID, StudentID, Name, Department string
	Eligible                        bool
	CreatedAt                       time.Time
}
type Vote struct {
	ID, ClubID, VoterID, RequestID string
	CreatedAt                      time.Time
}
type Review struct {
	ID, ClubID, ReviewerID, Decision, Comment string
	CreatedAt                                 time.Time
}
type RankingEntry struct {
	ClubID, Name, Category string
	Votes                  int
	Position               int
}

func ValidateClub(c Club) error {
	if strings.TrimSpace(c.ID) == "" || strings.TrimSpace(c.Name) == "" {
		return errors.New("club id and name required")
	}
	if c.VoteCount < 0 {
		return errors.New("negative votes")
	}
	return nil
}
func ValidateVoter(v Voter) error {
	if v.ID == "" || v.StudentID == "" {
		return errors.New("voter identity required")
	}
	if v.Eligible && v.Department == "" {
		return errors.New("department required")
	}
	return nil
}
func ValidateVote(v Vote) error {
	if v.ID == "" || v.ClubID == "" || v.VoterID == "" {
		return errors.New("vote fields required")
	}
	return nil
}
func ValidateReview(r Review) error {
	if r.ID == "" || r.ClubID == "" || r.ReviewerID == "" {
		return errors.New("review fields required")
	}
	if r.Decision != "approved" && r.Decision != "rejected" {
		return errors.New("invalid decision")
	}
	return nil
}
func IsPublished(c Club) bool { return c.Status == "published" }
func NormalizeStatus(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return "draft"
	}
	return s
}
func NewClub(id, name, cat, desc string) Club {
	return Club{ID: id, Name: name, Category: cat, Description: desc, Status: "draft", CreatedAt: time.Now().UTC()}
}
func NewVoter(id, student, name, dept string, eligible bool) Voter {
	return Voter{ID: id, StudentID: student, Name: name, Department: dept, Eligible: eligible, CreatedAt: time.Now().UTC()}
}
func NewVote(id, club, voter, req string) Vote {
	return Vote{ID: id, ClubID: club, VoterID: voter, RequestID: req, CreatedAt: time.Now().UTC()}
}
func NewReview(id, club, reviewer, decision, comment string) Review {
	return Review{ID: id, ClubID: club, ReviewerID: reviewer, Decision: decision, Comment: comment, CreatedAt: time.Now().UTC()}
}
