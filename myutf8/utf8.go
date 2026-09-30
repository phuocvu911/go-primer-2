package myutf8

func ValidRune(r rune) bool
func RuneLen(r rune) int
func EncodeRune(p []byte, r rune) (int, bool)
func AppendRune(p []byte, r rune) []byte
