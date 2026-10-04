package myunicode

// IsDigit return true if the rune is a decimal digit.
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsLetter return true if the rune is a letter.
func IsLetter(r rune) bool {
	return IsUpper(r) || IsLower(r)
}

// IsSpace return true if the rune is a whitespace character (ASCII only)
func IsSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\v' || r == '\f' || r == '\r' || r == '\n'
}

// IsUpper return true if the rune is an uppercase letter
func IsUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

// IsLower return true if the rune is a lowercase letter
func IsLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

// IsPunct return true if the rune is a punctuation character (ASCII only)
func IsPunct(r rune) bool {
	return r > 0x20 && r < 0x7F && !IsLetter(r) && !IsDigit(r)
}

// IsControl return true if the rune is a control character (ASCII only)
func IsControl(r rune) bool {
	return (r >= 0 && r < 0x20) || r == 0x7F //DEL
}

// ToUpper return the uppercase version of the rune
func ToUpper(r rune) rune {
	if IsLower(r) {
		r -= 'a' - 'A'
	}
	return r
}

// ToLower return the lowercase version of the rune
func ToLower(r rune) rune {
	if IsUpper(r) {
		r += 'a' - 'A'
	}
	return r
}
