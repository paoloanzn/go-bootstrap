package format

import (
	"regexp"

	"github.com/paoloanzn/go-bootstrap/config"
)

func DefaultWildCards() map[string]string {
	w := make(map[string]string)

	w["<main_package>"] = config.Cfg.ProjectName

	return w
}

func MatchWildCards(s string) string {
	defaults := DefaultWildCards()
	re := regexp.MustCompile(`<[^>]+>`) // match <string> pattern

	replaced := s // copy

	allMatches := re.FindAllStringSubmatch(s, -1)
	for _, m := range allMatches {
		if len(m) > 0 {
			key := m[0]
			value, exists := defaults[key]
			if exists {
				replaced = re.ReplaceAllString(replaced, value)
			}
		}
	}

	return replaced
}
