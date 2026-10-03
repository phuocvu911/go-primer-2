package myutf8

// ValidRune checks r is a valid Unicode code point.
func ValidRune(r rune) bool {
	// 1. Check if it is in the surrogate half range
	if r >= 0xD800 && r <= 0xDFFF {
		return false
	}

	// 2. Check if it exceeds the maximum Unicode code point
	if r > 0x10FFFF {
		return false
	}

	// 3. Optional: Filter out negative values if your input type allows them
	if r < 0 {
		return false
	}

	return true
}

const (
	maxRune1 = 0x7f
	maxRune2 = 0x7ff
	maxRune3 = 0xffff
	maxrune4 = 0x10ffff
)

// RuneLen returns the number of bytes required to encode the rune r in UTF-8.
//
// Bytes	Code point range	Payload bits	Byte pattern
// 1		U+0000 – U+007F		7				0xxxxxxx
// 2		U+0080 – U+07FF		11				110xxxxx 10xxxxxx
// 3		U+0800 – U+FFFF		16				1110xxxx 10xxxxxx 10xxxxxx
// 4		U+10000 – U+10FFFF	21				11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
func RuneLen(r rune) int {
	if !ValidRune(r) {
		return -1
	}
	switch {
	case r <= maxRune1:
		return 1
	case r <= maxRune2:
		return 2
	case r <= maxRune3:
		return 3
	case r <= maxrune4:
		return 4
	}
	return -1
}

// EncodeRune encodes the rune r into p and returns the number of bytes written.
func EncodeRune(p []byte, r rune) (int, bool) {

}

// // AppendRune appends the UTF-8 encoding of the rune r to p and returns the extended slice.
// func AppendRune(p []byte, r rune) []byte {

// }
