package api

import (
	"campusawards/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"
)

type ClubReader interface {
	ListClubs() ([]domain.Club, error)
	GetClub(string) (domain.Club, error)
}

func ClubsHandler(reader ClubReader) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "method not allowed", 405)
			return
		}
		clubs, e := reader.ListClubs()
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(clubs)
	})
}
func ClubHandler(reader ClubReader) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "id required", 400)
			return
		}
		c, e := reader.GetClub(id)
		if e != nil {
			http.Error(w, "not found", 404)
			return
		}
		json.NewEncoder(w).Encode(c)
	})
}
func ParseLimit(r *http.Request, def int) int {
	v, e := strconv.Atoi(r.URL.Query().Get("limit"))
	if e != nil || v < 1 {
		return def
	}
	if v > 100 {
		return 100
	}
	return v
}
func WriteJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func ErrorJSON(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	WriteJSON(w, map[string]string{"error": msg})
}
