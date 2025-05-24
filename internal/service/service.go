package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Conv(data string) (string, error) {

	if len(data) == 0 {
		return "", errors.New("mistake in data")
	}

	checkSym := func(sym rune) bool {
		return sym != '-' && sym != '.' && !unicode.IsSpace(sym)
	}

	isText := strings.ContainsFunc(data, checkSym)

	if isText {
		return morse.ToMorse(data), nil
	}
	return morse.ToText(data), nil
}
