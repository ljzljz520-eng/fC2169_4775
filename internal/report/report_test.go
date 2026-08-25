package report

import (
	"campusawards/internal/domain"
	"testing"
)

func TestReportExport(t *testing.T) {
	s := Export([]domain.RankingEntry{{Position: 1, Name: "A", Category: "x", Votes: 2}}, "评选")
	if s == "" {
		t.Fatal("empty")
	}
}
