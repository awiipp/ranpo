package cmd

import "strings"

func parseHeaders(raw []string) map[string]string {
	h := map[string]string{}
	for _, s := range raw {
		idx := strings.Index(s, ":")
		if idx > 0 {
			h[strings.TrimSpace(s[:idx])] = strings.TrimSpace(s[idx+1:])
		}
	}

	return h
}
