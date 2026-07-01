package pathutil

import "testing"

func TestExpandTilde(t *testing.T) {
	home := "/home/tester"
	t.Setenv("HOME", home)

	tests := []struct {
		input    string
		expected string
	}{
		{"~/src", home + "/src"},
		{"~", home},
		{"/absolute/path", "/absolute/path"},
		{"relative/path", "relative/path"},
		{"~user/path", "~user/path"}, // only bare ~ is expanded
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ExpandTilde(tt.input); got != tt.expected {
				t.Errorf("ExpandTilde(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestExpandTildeNoHome(t *testing.T) {
	t.Setenv("HOME", "")
	if got := ExpandTilde("~/src"); got != "~/src" {
		t.Errorf("ExpandTilde with no HOME = %q, want %q", got, "~/src")
	}
}
