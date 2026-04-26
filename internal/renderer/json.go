package renderer

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
)

var trailingComma = regexp.MustCompile(`,\s*([\]}])`)

// PrettyJSON formats and colorizes raw JSON bytes.
func PrettyJSON(data []byte) string {
	clean := trailingComma.ReplaceAll(bytes.TrimSpace(data), []byte("${1}"))

	var buf bytes.Buffer
	if err := json.Indent(&buf, clean, "", "  "); err != nil {
		return string(data)
	}

	return colorizeJSON(strings.TrimRight(buf.String(), "\n"))
}