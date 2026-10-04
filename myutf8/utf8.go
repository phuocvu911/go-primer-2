package myutf8

const (
	maxRune1  = 0x7f
	maxRune2  = 0x7ff
	maxRune3  = 0xffff
	maxrune4  = 0x10ffff
	RuneError = '\uFFFD'
)

// ValidRune checks r is a valid Unicode code point.
func ValidRune(r rune) bool {
	// the surrogate half range
	if r >= 0xD800 && r <= 0xDFFF {
		return false
	}

	//exceeds the maximum Unicode code point or negative
	if r > 0x10FFFF || r < 0 {
		return false
	}

	return true
}

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
// p has to be large enough to contain r
func EncodeRune(p []byte, r rune) (int, bool) {
	//check valid r
	if !ValidRune(r) {
		r = RuneError
	}

	//valid p
	lenRune := RuneLen(r)
	if lenRune > len(p) {
		return 0, false
	}

	switch lenRune {
	case 1:
		p[0] = byte(r)
	case 2:
		p[0] = 0xC0 | byte(r>>6)   //1100 0000
		p[1] = 0x80 | byte(r&0x3F) // 1000 0000 - 00111111
	case 3:
		p[0] = 0xE0 | byte(r>>12)
		p[1] = 0x80 | byte(r>>6&0x3F)
		p[2] = 0x80 | byte(r&0x3F)
	case 4:
		p[0] = 0xF0 | byte(r>>18)
		p[1] = 0x80 | byte(r>>12&0x3F)
		p[2] = 0x80 | byte(r>>6&0x3F)
		p[3] = 0x80 | byte(r&0x3F)
	}
	return lenRune, true
}

// AppendRune appends the UTF-8 encoding of the rune r to p and returns the extended slice.
func AppendRune(p []byte, r rune) []byte {
	if !ValidRune(r) {
		r = RuneError
	}
	tail := make([]byte, RuneLen(r))
	_, _ = EncodeRune(tail, r)
	p = append(p, tail...)
	return p
}
