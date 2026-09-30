package myunicode

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsLetter(r rune) bool {
	return IsUpper(r) || IsLower(r)
}

func IsSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\v' || r == '\f' || r == '\r' || r == '\n'
}

func IsUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func IsLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func IsPunct(r rune) bool {
	return r > 0x20 && r < 0x7F && !IsLetter(r) && !IsDigit(r)
}

func IsControl(r rune) bool {
	return (r >= 0 && r < 0x20) || r == 0x7F //DEL
}

func ToUpper(r rune) rune {
	if IsLower(r) {
		r -= 'a' - 'A'
	}
	return r
}

func ToLower(r rune) rune {
	if IsUpper(r) {
		r += 'a' - 'A'
	}
	return r
}
