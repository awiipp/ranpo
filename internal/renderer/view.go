package renderer

import "strings"

// HeadersView formats response headers for display.
func HeadersView(headers map[string][]string) string {
	var sb strings.Builder
	for k, vs := range headers {
		sb.WriteString(labelStyle.Render(k))
		sb.WriteString(dimStyle.Render(": "))
		sb.WriteString(strings.Join(vs, ", "))
		sb.WriteString("\n")
	}
	return sb.String()
}

func indentBlock(s string, indent string) string {
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		lines[i] = indent + line
	}

	return strings.Join(lines, "\n")
}
