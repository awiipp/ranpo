package client

import "strings"

// Resolve replaces {{KEY}} placeholders with values from vars.
func Resolve(s string, vars map[string]string) string {
	for k, v := range vars {
		s = strings.ReplaceAll(s, "{{"+k+"}}", v)
	}

	return s
}
