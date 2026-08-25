package api

import (
	"campusawards/internal/ranking"
	"net/http/httptest"
	"testing"
)

func TestAPIHealth(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	NewServer(&ranking.Service{}).Handler().ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Fatal(rr.Code)
	}
}
