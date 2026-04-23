package store

import "strings"

func sanitize(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "_"))
}
