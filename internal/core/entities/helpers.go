package entities

import (
	"fmt"
	"strings"
)

func ConvertToWildcards(strs []string) []Wildcard {
	wildcards := make([]Wildcard, len(strs))
	for i, s := range strs {
		if !strings.HasPrefix(s, "/") {
			s = fmt.Sprintf("/%s", s)
		}
		wildcards[i] = Wildcard(s)
	}
	return wildcards
}
