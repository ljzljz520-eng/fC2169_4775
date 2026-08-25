package domain

import "strings"

func ClubLabel(c Club) string { return strings.TrimSpace(c.Category) + "/" + strings.TrimSpace(c.Name) }
func StatusDescription(status string) string {
	switch NormalizeStatus(status) {
	case "draft":
		return "待审核"
	case "published":
		return "已发布"
	case "rejected":
		return "未通过"
	case "archived":
		return "已归档"
	default:
		return "未知"
	}
}
func ReviewDecision(decision string) bool { return decision == "approved" || decision == "rejected" }
func VoteWeight(v Vote) int {
	if v.RequestID == "" {
		return 0
	}
	return 1
}
