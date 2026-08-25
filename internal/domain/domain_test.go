package domain

import "testing"

func TestClubValidation(t *testing.T) {
	if ValidateClub(NewClub("", "x", "c", "")) == nil {
		t.Fatal("expected error")
	}
	if ValidateClub(NewClub("1", "x", "c", "")) != nil {
		t.Fatal("valid club rejected")
	}
}
