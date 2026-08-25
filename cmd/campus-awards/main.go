package main

import (
	"campusawards/internal/api"
	"campusawards/internal/ranking"
	"campusawards/internal/store"
	"flag"
	"fmt"
	"log"
)

func main() {
	path := flag.String("db", "campus-awards.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	s, e := store.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	r := ranking.NewService(s)
	fmt.Printf("校园十佳社团评选服务 listening on %s\n", *addr)
	log.Fatal(api.NewServer(r).Serve(*addr))
}
