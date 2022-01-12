package threp

import "strings"

func trimNull(s string) string {
	return strings.TrimRightFunc(s, func(r rune) bool { return r == 0 })
}

func safeIndex(arr []string, index byte) string {
	if int(index) >= len(arr) {
		return "Unknown"
	}
	return arr[index]
}
