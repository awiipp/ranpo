package renderer

import "bytes"

func colorizeJSON(s string) string {
	var out bytes.Buffer
	runes := []rune(s)
	i := 0
	isKey := true // first string in an object is always a key

	for i < len(runes) {
		ch := runes[i]

		switch {
		case ch == '"':
			// Collect entire string token including quotes
			start := i
			i++

			for i < len(runes) {
				if runes[i] == '"' && runes[i-1] != '\\' {
					break
				}
				i++
			}

			token := string(runes[start : i+1])
			if isKey {
				out.WriteString(keyStyle.Render(token))
				isKey = false
			} else {
				out.WriteString(strStyle.Render(token))
			}

			i++
			continue

		case ch == ':':
			out.WriteString(punctStyle.Render(":"))

		case ch == '{' || ch == '[':
			isKey = (ch == '{')
			out.WriteString(punctStyle.Render(string(ch)))

		case ch == '}' || ch == ']':
			out.WriteString(punctStyle.Render(string(ch)))

		case ch == ',':
			out.WriteString(punctStyle.Render(","))
			// After a comma in an object context, next string is a key
			// This detected by looking ahead for the next "
			for j := i + 1; j < len(runes); j++ {
				if runes[j] == '"' {
					isKey = true
					break
				}
				if runes[j] != ' ' && runes[j] != '\n' && runes[j] != '\r' && runes[j] != '\t' {
					break
				}
			}

		case i+4 <= len(runes) && string(runes[i:i+4]) == "true":
			out.WriteString(boolStyle.Render("true"))
			i += 4
			continue

		case i+5 <= len(runes) && string(runes[i:i+5]) == "false":
			out.WriteString(boolStyle.Render("false"))
			i += 5
			continue

		case i+4 <= len(runes) && string(runes[i:i+4]) == "null":
			out.WriteString(nullStyle.Render("null"))
			i += 4
			continue

		case (ch >= '0' && ch <= '9') || ch == '-':
			start := i
			for i < len(runes) {
				c := runes[i]
				if (c >= '0' && c <= '9') || c == '.' || c == '-' || c == 'e' || c == 'E' || c == '+' {
					i++
				} else {
					break
				}
			}
			out.WriteString(numStyle.Render(string(runes[start:i])))
			continue

		default:
			out.WriteRune(ch)
		}

		i++
	}

	return out.String()
}
