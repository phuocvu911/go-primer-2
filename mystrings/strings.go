package mystrings

import (
	"goprimer/myunicode"
)

// Index return the byte index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
	n, m := len(s), len(substr)
	//substr is 0
	if m == 0 {
		return 0
	}

	//substr is longer than s
	if m > n {
		return -1
	}

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && s[i+j] == substr[j] {
			j++
		}
		if j == m {
			return i
		}
	}
	return -1
}

// HasPrefix return true if s begins with prefix.
func HasPrefix(s, prefix string) bool {
	return Index(s, prefix) == 0
}

// Cut splits s around the first instance of sep, returning the text before and after sep.
func Cut(s, sep string) (before, after string, found bool) {
	idx := Index(s, sep)
	if idx == -1 {
		return s, "", false
	}
	return s[:idx], s[idx+len(sep):], true
}

// Count return the number of non-overlapping instances of substr in s.
func Count(s, substr string) int {
	cnt := 0
	if len(substr) == 0 {
		return len([]rune(s)) + 1
	}
	for {
		idx := Index(s, substr)
		if idx == -1 {
			break
		} else {
			_, s, _ = Cut(s, substr)
			cnt++
		}
	}
	return cnt
}

// IndexRune return the first byte index of rune r in s.
func IndexRune(s string, r rune) int {
	for i, ch := range s {
		if ch == r {
			return i
		}
	}
	return -1
}

// TrimSpace return a slice of the string s, with all leading and trailing white space removed (ASCII only).
func TrimSpace(s string) string {
	head := 0
	for i := 0; i < len(s); i++ {
		if !myunicode.IsSpace(rune(s[i])) {
			break
		}
		head = i + 1
	}
	tail := len(s)
	for j := tail - 1; j > head; j-- {
		if !myunicode.IsSpace(rune(s[j])) {
			break
		}
		tail = j
	}
	return s[head:tail]
}

// ToLower return a copy of the string s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
	res := ""
	for _, r := range s {
		if myunicode.IsUpper(r) {
			res += string(myunicode.ToLower(r))
		} else {
			res += string(r)
		}
	}
	return res
}

// Repeat return a copy of the string s concatenated count times.
func Repeat(s string, count int) (string, bool) {
	if count < 0 {
		return "", false
	}

	res := ""
	for range count {
		res += s
	}
	return res, true
}
