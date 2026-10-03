package mystrings

import (
	"testing"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		substr string
		want   int
	}{
		{"found", "hello world", "world", 6},
		{"missing", "hello", "x", -1},
		{"empty substring", "hello", "", 0},
		{"unicode byte index", "café", "é", 3},
		{"empty string", "", "a", -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Index(test.s, test.substr); got != test.want {
				t.Errorf("Index(%q, %q) = %d, want %d", test.s, test.substr, got, test.want)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		prefix string
		want   bool
	}{
		{"matching", "gopher", "go", true},
		{"not matching", "gopher", "ph", false},
		{"empty prefix", "gopher", "", true},
		{"longer prefix", "go", "gopher", false},
		{"unicode", "éclair", "é", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := HasPrefix(test.s, test.prefix); got != test.want {
				t.Errorf("HasPrefix(%q, %q) = %t, want %t", test.s, test.prefix, got, test.want)
			}
		})
	}
}

func TestCut(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		sep        string
		wantBefore string
		wantAfter  string
		wantFound  bool
	}{
		{"found", "name=value", "=", "name", "value", true},
		{"missing", "name", "=", "name", "", false},
		{"first separator", "a:b:c", ":", "a", "b:c", true},
		{"empty separator", "abc", "", "", "abc", true},
		{"separator at start", ":value", ":", "", "value", true},
		{"separator at end", "name:", ":", "name", "", true},
		{"unicode separator", "café", "é", "caf", "", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before, after, found := Cut(test.s, test.sep)
			if before != test.wantBefore || after != test.wantAfter || found != test.wantFound {
				t.Errorf("Cut(%q, %q) = (%q, %q, %t), want (%q, %q, %t)", test.s, test.sep, before, after, found, test.wantBefore, test.wantAfter, test.wantFound)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		substr string
		want   int
	}{
		{"repeated", "banana", "an", 2},
		{"overlapping matches are not counted", "aaa", "aa", 1},
		{"empty substring", "abc", "", 4},
		{"missing", "abc", "x", 0},
		{"unicode", "ééé", "é", 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Count(test.s, test.substr); got != test.want {
				t.Errorf("Count(%q, %q) = %d, want %d", test.s, test.substr, got, test.want)
			}
		})
	}
}

func TestIndexRune(t *testing.T) {
	tests := []struct {
		name string
		s    string
		r    rune
		want int
	}{
		{"ASCII", "hello", 'e', 1},
		{"unicode byte index", "café", 'é', 3},
		{"missing", "hello", 'x', -1},
		{"NUL", "a\x00b", 0, 1},
		{"empty string", "", 'a', -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IndexRune(test.s, test.r); got != test.want {
				t.Errorf("IndexRune(%q, %U) = %d, want %d", test.s, test.r, got, test.want)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"  hello  ", "hello"},
		{"\t\nhello\r\n", "hello"},
		{"\u00a0hello\u00a0", "hello"},
		{"", ""},
		{"already trimmed", "already trimmed"},
		{" \t\n", ""},
	}
	for _, test := range tests {
		if got := TrimSpace(test.input); got != test.want {
			t.Errorf("TrimSpace(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Hello, WORLD!", "hello, world!"},
		{"Already lower", "already lower"},
		{"ÉCLAIR", "éclair"},
		{"123!", "123!"},
		{"", ""},
	}
	for _, test := range tests {
		if got := ToLower(test.input); got != test.want {
			t.Errorf("ToLower(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		count int
		want  string
		ok    bool
	}{
		{"positive count", "go", 3, "gogogo", true},
		{"zero count", "go", 0, "", true},
		{"empty string", "", 5, "", true},
		{"negative count", "go", -1, "", false},
		{"unicode", "é", 2, "éé", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := Repeat(test.input, test.count)
			if got != test.want || ok != test.ok {
				t.Errorf("Repeat(%q, %d) = (%q, %t), want (%q, %t)", test.input, test.count, got, ok, test.want, test.ok)
			}
		})
	}
}
