package common

import "strings"

// helper: check substring ignore case
func ContainsIgnoreCase(str, substr string) bool {
	strLower := strings.ToLower(str)
	subLower := strings.ToLower(substr)
	return strings.Contains(strLower, subLower)
}
