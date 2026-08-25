package ranking

import (
	"campusawards/internal/domain"
	"fmt"
	"strings"
)

func Format(entries []domain.RankingEntry) string {
	var b strings.Builder
	for _, e := range entries {
		fmt.Fprintf(&b, "%d. %s [%s] %d\n", e.Position, e.Name, e.Category, e.Votes)
	}
	return b.String()
}
func Medal(position int) string {
	switch position {
	case 1:
		return "gold"
	case 2:
		return "silver"
	case 3:
		return "bronze"
	default:
		return "finalist"
	}
}
func IsTop(position int) bool { return position >= 1 && position <= 10 }
