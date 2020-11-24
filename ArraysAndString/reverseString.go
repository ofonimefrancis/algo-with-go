package ArraysAndString

import (
	"unicode"
	"unicode/utf8"
)

// func main() {
// 	fmt.Println(reverse("string"))
// }

func reverse(s string) string {
	runes := []rune(s)
	ret := ""
	for i := len(runes) - 1; i >= 0; i-- {
		ret = ret + string(runes[i]) //Inefficient: Strings are Immutable
	}
	return ret
}

func reverseString(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

// no encoding
func reverseBytes(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

// reverseCodePoints interprets its argument as UTF-8 and ignores bytes
// that do not form valid UTF-8.  return value is UTF-8.
func reverseCodePoints(s string) string {
	r := make([]rune, len(s))
	start := len(s)
	for _, c := range s {
		// quietly skip invalid UTF-8
		if c != utf8.RuneError {
			start--
			r[start] = c
		}
	}
	return string(r[start:])
}

// reversePreservingCombiningCharacters interprets its argument as UTF-8
// and ignores bytes that do not form valid UTF-8.  return value is UTF-8.
func reversePreservingCombiningCharacters(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		// quietly skip invalid UTF-8
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return (string(r[start:]))
}
