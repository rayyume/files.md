package txt

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func I64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

func Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

// Substr isn't multi-Unicode-codepoint aware, like specifying skintone or
// gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)
	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func Emoji(emoji, str string) string {
	return fmt.Sprintf("%s %s", emoji, str)
}

func NormNewLines(text string) string {
	text = strings.Replace(text, "\\r\\n", "\n", -1)
	return strings.Replace(text, "\\n\\r", "\n", -1)
}

func SplitTextIntoChunks(text string, maxLen int) []string {
	var chunks []string

	for utf8.RuneCountInString(text) > maxLen {
		// Get the substring of the first maxLen runes
		runes := []rune(text)
		subStr := string(runes[:maxLen])

		// Find the last newline in the substring
		splitIndex := strings.LastIndex(subStr, "\n")
		if splitIndex == -1 {
			// No newline found, find the last space
			splitIndex = strings.LastIndex(subStr, " ")
			if splitIndex == -1 {
				// No space found either, split at maxLen
				splitIndex = maxLen
			}
		} else {
			// Adjust the split index to the rune count
			splitIndex = utf8.RuneCountInString(subStr[:splitIndex])
		}

		chunks = append(chunks, string(runes[:splitIndex]))
		text = string(runes[splitIndex:])
	}
	chunks = append(chunks, text)

	return chunks
}
