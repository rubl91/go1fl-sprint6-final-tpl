package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseConverter(data []byte) string {
	s := string(data)
	r := strings.NewReplacer(".", "", "-", "", " ", "")
	replacedString := r.Replace(s)
	if replacedString == "" {
		return morse.ToText(s)
	} else {
		return morse.ToMorse(s)
	}
}
