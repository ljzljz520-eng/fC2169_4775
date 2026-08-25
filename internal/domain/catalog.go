package domain

import (
	"sort"
	"strings"
	"time"
)

type Catalog struct {
	Clubs     map[string]Club
	UpdatedAt time.Time
}

func NewCatalog() Catalog { return Catalog{Clubs: map[string]Club{}, UpdatedAt: time.Now().UTC()} }
func (c *Catalog) Add(club Club) error {
	if e := ValidateClub(club); e != nil {
		return e
	}
	if c.Clubs == nil {
		c.Clubs = map[string]Club{}
	}
	c.Clubs[club.ID] = club
	c.UpdatedAt = time.Now().UTC()
	return nil
}
func (c *Catalog) Remove(id string) bool {
	if _, ok := c.Clubs[id]; !ok {
		return false
	}
	delete(c.Clubs, id)
	c.UpdatedAt = time.Now().UTC()
	return true
}
func (c Catalog) Get(id string) (Club, bool) { v, ok := c.Clubs[id]; return v, ok }
func (c Catalog) All() []Club {
	out := make([]Club, 0, len(c.Clubs))
	for _, v := range c.Clubs {
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
func (c Catalog) ByStatus(status string) []Club {
	out := []Club{}
	for _, v := range c.Clubs {
		if NormalizeStatus(v.Status) == NormalizeStatus(status) {
			out = append(out, v)
		}
	}
	return out
}
func (c Catalog) ByCategory(category string) []Club {
	out := []Club{}
	for _, v := range c.Clubs {
		if strings.EqualFold(v.Category, category) {
			out = append(out, v)
		}
	}
	return out
}
func (c Catalog) Count() int { return len(c.Clubs) }
func (c *Catalog) Publish(id string) bool {
	v, ok := c.Clubs[id]
	if !ok {
		return false
	}
	v.Status = "published"
	c.Clubs[id] = v
	return true
}
func (c *Catalog) Archive(id string) bool {
	v, ok := c.Clubs[id]
	if !ok || v.Status != "published" {
		return false
	}
	v.Status = "archived"
	c.Clubs[id] = v
	return true
}
func MergeCatalog(a, b Catalog) Catalog {
	out := NewCatalog()
	for _, v := range a.Clubs {
		out.Clubs[v.ID] = v
	}
	for _, v := range b.Clubs {
		out.Clubs[v.ID] = v
	}
	return out
}
func (c Catalog) Snapshot() []RankingEntry {
	out := []RankingEntry{}
	for _, v := range c.Clubs {
		if IsPublished(v) {
			out = append(out, RankingEntry{ClubID: v.ID, Name: v.Name, Category: v.Category, Votes: v.VoteCount})
		}
	}
	SortRankings(out)
	return out
}
