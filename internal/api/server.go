package api

import (
	"campusawards/internal/ranking"
	"encoding/json"
	"net/http"
)

type Server struct{ Ranking *ranking.Service }

func NewServer(r *ranking.Service) *Server { return &Server{Ranking: r} }
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte("ok")) })
	mux.HandleFunc("/ranking", s.ranking)
	return mux
}
func (s *Server) ranking(w http.ResponseWriter, r *http.Request) {
	entries, e := s.Ranking.TopTen()
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
func (s *Server) Serve(addr string) error { return http.ListenAndServe(addr, s.Handler()) }
