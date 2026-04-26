package renderer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

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
	var buf bytes.Buffer
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		return string(data)
	}

	return colorizeJSON(buf.String())
}

// RenderResponse builds the complete formatted response string for CLI output
func RenderResponse(code int, status string, body []byte, duration fmt.Stringer) string {
	divider := dimStyle.Render(strings.Repeat("─", 56))
	statusLine := StatusLine(code, status, duration, len(body))

	return fmt.Sprintf("\n%s\n%s\n\n%s\n", statusLine, divider, PrettyJSON(body))
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
