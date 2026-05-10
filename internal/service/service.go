package service

import "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"

func Converter(content string) string {
	isMorse := true
	for _, char := range content {
		if char != rune('.') && char != rune('-') && char != rune(' ') {
			isMorse = false
			break
		}
	}
	if isMorse {
		return morse.ToText(content)
	}
	return morse.ToMorse(content)
}
