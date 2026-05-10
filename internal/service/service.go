package service

import "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"

// funcrion for distinction morse and text content
func Converter(content string) string {
	// bool flag
	isMorse := true
	for _, char := range content {
		// checking every char for equaling morse symbols
		if char != rune('.') && char != rune('-') && char != rune(' ') {
			// argument is not a morse
			isMorse = false
			break
		}
	}
	// different methods depending on data types
	if isMorse {
		return morse.ToText(content)
	}
	return morse.ToMorse(content)
}
