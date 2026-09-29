package myunicode

func IsDigit(r rune) bool {
	if r >= 0x80 {
		return false
	}
}
func IsLetter(r rune) bool
func IsSpace(r rune) bool
func IsUpper(r rune) bool
func IsLower(r rune) bool
func IsPunct(r rune) bool
func IsControl(r rune) bool
func ToUpper(r rune) rune
func ToLower(r rune) rune
