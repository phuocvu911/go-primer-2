package myunicode

import (
	"testing"
	std "unicode"
)

func TestASCIIPredicatesMatchUnicode(t *testing.T) {
	tests := []struct {
		name string
		got  func(rune) bool
		want func(rune) bool
	}{
		{"IsDigit", IsDigit, std.IsDigit},
		{"IsLetter", IsLetter, std.IsLetter},
		{"IsSpace", IsSpace, std.IsSpace},
		{"IsUpper", IsUpper, std.IsUpper},
		{"IsLower", IsLower, std.IsLower},
		{"IsControl", IsControl, std.IsControl},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for r := rune(0); r < 0x80; r++ {
				if got, want := test.got(r), test.want(r); got != want {
					t.Errorf("%s(%U) = %t, want %t", test.name, r, got, want)
				}
			}
		})
	}
}

func TestASCIIPunctuation(t *testing.T) {
	tests := []struct {
		r    rune
		want bool
	}{
		{'!', true},
		{'~', true},
		{'0', false},
		{'A', false},
		{'a', false},
		{' ', false},
		{'\t', false},
		{0x7f, false},
	}

	for _, test := range tests {
		if got := IsPunct(test.r); got != test.want {
			t.Errorf("IsPunct(%U) = %t, want %t", test.r, got, test.want)
		}
	}
}

func TestASCIICaseConversionMatchesUnicode(t *testing.T) {
	tests := []struct {
		name string
		got  func(rune) rune
		want func(rune) rune
	}{
		{"ToUpper", ToUpper, std.ToUpper},
		{"ToLower", ToLower, std.ToLower},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for r := rune(0); r < 0x80; r++ {
				if got, want := test.got(r), test.want(r); got != want {
					t.Errorf("%s(%U) = %U, want %U", test.name, r, got, want)
				}
			}
		})
	}
}

func TestNonASCIIRunes(t *testing.T) {
	predicates := []struct {
		name string
		fn   func(rune) bool
	}{
		{"IsDigit", IsDigit},
		{"IsLetter", IsLetter},
		{"IsSpace", IsSpace},
		{"IsUpper", IsUpper},
		{"IsLower", IsLower},
		{"IsPunct", IsPunct},
		{"IsControl", IsControl},
	}
	conversion := []struct {
		name string
		fn   func(rune) rune
	}{
		{"ToUpper", ToUpper},
		{"ToLower", ToLower},
	}
	runes := []rune{-1, -0x80, 0x80, 0x85, 0xE9, 0x1F600, 0x10FFFF}

	for _, test := range predicates {
		t.Run(test.name, func(t *testing.T) {
			for _, r := range runes {
				if got := test.fn(r); got {
					t.Errorf("%s(%U) = true, want false", test.name, r)
				}
			}
		})
	}

	for _, test := range conversion {
		t.Run(test.name, func(t *testing.T) {
			for _, r := range runes {
				if got := test.fn(r); got != r {
					t.Errorf("%s(%U) = %U, want unchanged", test.name, r, got)
				}
			}
		})
	}
}
