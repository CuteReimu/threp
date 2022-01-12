package threp

import "strings"

func trim(s string) string {
	s = trimNull(s)
	s = strings.TrimSpace(s)
	return trimNull(s)
}

func getValue(key string, line string) string {
	line = trim(line)
	if strings.Index(line, key) != 0 {
		return ""
	}
	return line[len(key):]
}

func trimNull(s string) string {
	return strings.TrimRightFunc(s, func(r rune) bool { return r == 0 })
}

func safeIndex(arr []string, index byte) string {
	if int(index) >= len(arr) {
		return "Unknown"
	}
	return arr[index]
}
