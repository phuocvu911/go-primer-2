package myutf8

import (
	"bytes"
	"testing"
	std "unicode/utf8"
)

func TestValidRune(t *testing.T) {
	tests := []struct {
		r    rune
		want bool
	}{
		{-1, false},
		{0, true},
		{0x7F, true},
		{0x80, true},
		{0xD7FF, true},
		{0xD800, false},
		{0xDFFF, false},
		{0xE000, true},
		{std.MaxRune, true},
		{std.MaxRune + 1, false},
	}

	for _, test := range tests {
		if got := ValidRune(test.r); got != test.want {
			t.Errorf("ValidRune(%U) = %t, want %t", test.r, got, test.want)
		}
	}
}

func TestRuneLen(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want int
	}{
		{"negative", -1, -1},
		{"null", 0, 1},
		{"ascii", 0x7F, 1},
		{"two byte", 0x80, 2},
		{"two byte max", 0x7FF, 2},
		{"three byte", 0x800, 3},
		{"three byte max", 0xFFFF, 3},
		{"four byte", 0x10000, 4},
		{"max rune", std.MaxRune, 4},
		{"surrogate", 0xD800, -1},
		{"too large", std.MaxRune + 1, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := RuneLen(test.r); got != test.want {
				t.Errorf("RuneLen(%U) = %d, want %d", test.r, got, test.want)
			}
		})
	}
}

func TestEncodeRune(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want []byte
	}{
		{"ASCII", 'A', []byte("A")},
		{"two-byte", 0xA2, []byte{0xC2, 0xA2}},
		{"three-byte", 0x20AC, []byte{0xE2, 0x82, 0xAC}},
		{"four-byte", 0x1F600, []byte{0xF0, 0x9F, 0x98, 0x80}},
		{"maximum-rune", std.MaxRune, []byte{0xF4, 0x8F, 0xBF, 0xBF}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buffer := make([]byte, 4)
			n, ok := EncodeRune(buffer, test.r)
			if !ok {
				t.Fatalf("EncodeRune(%U) reported failure", test.r)
			}
			if n != len(test.want) {
				t.Errorf("EncodeRune(%U) wrote %d bytes, want %d", test.r, n, len(test.want))
			}
			if !bytes.Equal(buffer[:n], test.want) {
				t.Errorf("EncodeRune(%U) = % X, want % X", test.r, buffer[:n], test.want)
			}
		})
	}
}

func TestEncodeRuneInvalidRune(t *testing.T) {
	for _, r := range []rune{-1, 0xD800, std.MaxRune + 1} {
		buffer := bytes.Repeat([]byte{0xA5}, 4)
		n, ok := EncodeRune(buffer, r)
		if !ok || n != 3 {
			t.Errorf("EncodeRune(%U) = (%d, %t), want (3, true)", r, n, ok)
		}
		if !bytes.Equal(buffer, []byte{0xEF, 0xBF, 0xBD, 0xA5}) {
			t.Errorf("EncodeRune(%U) unexpected output", r)
		}
	}
}

func TestEncodeRuneRejectsShortBuffer(t *testing.T) {
	for _, test := range []struct {
		r    rune
		need int
	}{
		{'A', 1},
		{0xA2, 2},
		{0x20AC, 3},
		{0x1F600, 4},
	} {
		buffer := bytes.Repeat([]byte{0xA5}, test.need-1)
		n, ok := EncodeRune(buffer, test.r)
		if ok || n != 0 {
			t.Errorf("EncodeRune(%U) with buffer length %d = (%d, %t), want (0, false)", test.r, len(buffer), n, ok)
		}
		if !bytes.Equal(buffer, bytes.Repeat([]byte{0xA5}, len(buffer))) {
			t.Errorf("EncodeRune(%U) modified a short buffer on failure", test.r)
		}
	}
}

// func TestAppendRune(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		base []byte
// 		r    rune
// 		want []byte
// 	}{
// 		{"ASCII", []byte("prefix:"), 'A', []byte("prefix:A")},
// 		{"two-byte", []byte("prefix:"), 0xA2, []byte("prefix:\xC2\xA2")},
// 		{"three-byte", nil, 0x20AC, []byte{0xE2, 0x82, 0xAC}},
// 		{"four-byte", []byte{0x01}, 0x1F600, []byte{0x01, 0xF0, 0x9F, 0x98, 0x80}},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			if got := AppendRune(test.base, test.r); !bytes.Equal(got, test.want) {
// 				t.Errorf("AppendRune(% X, %U) = % X, want % X", test.base, test.r, got, test.want)
// 			}
// 		})
// 	}
// }

// func TestAppendRuneRejectsInvalidRune(t *testing.T) {
// 	base := []byte("prefix")
// 	for _, r := range []rune{-1, 0xD800, std.MaxRune + 1} {
// 		got := AppendRune(base, r)
// 		if !bytes.Equal(got, base) {
// 			t.Errorf("AppendRune(%q, %U) = % X, want unchanged % X", base, r, got, base)
// 		}
// 	}
// }
