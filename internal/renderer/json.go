package renderer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var trailingComma = regexp.MustCompile(`,\s*([\]}])`)

// StatusLine renders a colored status code + metadata line.
func StatusLine(code int, status string, duration fmt.Stringer, bodyLen int) string {
	var codeStr string
	switch {
	case code >= 200 && code < 300:
		codeStr = statusOKStyle.Render(fmt.Sprintf("%d", code))
	case code >= 300 && code < 400:
		codeStr = statusWarnStyle.Render(fmt.Sprintf("%d", code))
	default:
		codeStr = statusErrStyle.Render(fmt.Sprintf("%d", code))
	}

	return fmt.Sprintf(
		"  %s %s   %s   %s",
		codeStr,
		dimStyle.Render(status),
		dimStyle.Render(duration.String()),
		dimStyle.Render(fmt.Sprintf("%d bytes", bodyLen)),
	)
}

// PrettyJSON formats and colorizes raw JSON bytes.
func PrettyJSON(data []byte) string {
	clean := trailingComma.ReplaceAll(bytes.TrimSpace(data), []byte("${1}"))

	var buf bytes.Buffer
	if err := json.Indent(&buf, clean, "", "  "); err != nil {
		return string(data)
	}

	return colorizeJSON(strings.TrimRight(buf.String(), "\n"))
}

// RenderResponse builds the complete formatted response string for CLI output
func RenderResponse(code int, status string, body []byte, duration fmt.Stringer) string {
	divider := dimStyle.Render(strings.Repeat("─", 56))
	statusLine := StatusLine(code, status, duration, len(body))

	const bodyIndent = "  "
	formattedBody := indentBlock(PrettyJSON(body), bodyIndent)

	return fmt.Sprintf("\n%s\n%s\n\n%s\n", statusLine, divider, formattedBody)
}

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
