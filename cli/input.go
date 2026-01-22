package cli

import (
	"errors"
)

func ParseArgs(raw []string) (string, error) {
	if len(raw) != 2 {
		return "", errors.New("usage: github-activity \"username\"")
	}

	value := raw[1]
	return value, nil
}
