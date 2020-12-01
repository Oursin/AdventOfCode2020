package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFileLines(file string) ([]string, error) {
	var content []string
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	rd := bufio.NewReader(f)
	for err == nil {
		var s string
		s, err = rd.ReadString('\n')
		if err == nil {
			content = append(content, strings.TrimRight(s, "\n"))
		}
	}
	if err != io.EOF {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return content, nil
}