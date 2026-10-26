package service

import (
	"errors"
	"strings"

	"morze_converter/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("empty input")
	}

	if isMorse(data) {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil
}

func isMorse(s string) bool {
	for _, r := range s {
		switch r {
		case '.', '-', ' ', '\n', '\r', '\t':
		default:
			return false
		}
	}
	return strings.ContainsAny(s, ".-")
}
