package threp

import (
	"github.com/pkg/errors"
	"io"
	"strings"
)

func seek(reader io.Reader, offset int64) error {
	if offset == 0 {
		return nil
	}
	if r, ok := reader.(io.Seeker); ok {
		_, err := r.Seek(offset, io.SeekCurrent)
		return err
	}
	if offset < 0 {
		return errors.Errorf("cannot seek negative offset: %d", offset)
	}
	buf := make([]byte, offset)
	n, err := reader.Read(buf)
	if err != nil {
		return err
	}
	if int64(n) < offset {
		return io.EOF
	}
	return nil
}

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
