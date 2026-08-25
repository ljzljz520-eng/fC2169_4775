package review

import "campusawards/internal/domain"

func DecisionLabel(r domain.Review) string {
	if r.Decision == "approved" {
		return "通过"
	}
	return "退回"
}
func IsFinal(r domain.Review) bool         { return r.Decision == "approved" || r.Decision == "rejected" }
func RequiresComment(r domain.Review) bool { return r.Decision == "rejected" }
