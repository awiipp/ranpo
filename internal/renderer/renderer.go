package renderer

import (
	"fmt"
	"strings"
)

// RenderResponse builds the complete formatted response string for CLI output
func RenderResponse(code int, status string, body []byte, duration fmt.Stringer) string {
	divider := dimStyle.Render(strings.Repeat("─", 70))
	statusLine := StatusLine(code, status, duration, len(body))

	const bodyIndent = "  "
	formattedBody := indentBlock(PrettyJSON(body), bodyIndent)

	return fmt.Sprintf("\n%s\n%s\n\n%s\n", statusLine, divider, formattedBody)
}

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
