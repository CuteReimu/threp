package threp

import (
	"bufio"
	"io"
	"strings"

	"github.com/pkg/errors"
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
	n, err := io.ReadFull(reader, buf)
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

func linesToMap(fin io.Reader, separator string) (map[string]string, error) {
	reader := bufio.NewReader(fin)
	ret := make(map[string]string)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return nil, err
		}
		arr := strings.SplitN(string(line), separator, 2)
		if len(arr) == 2 {
			ret[trim(arr[0])] = trim(arr[1])
		}
		if err == io.EOF {
			break
		}
	}
	return ret, nil
}
