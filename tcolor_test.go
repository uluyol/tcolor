package tcolor

import "testing"

func TestHasColor(t *testing.T) {
	var tests = []struct {
		Name     string
		HasColor bool
	}{
		{"xterm", true},
		{"dumb", false},
		{"xterm-256color", true},
		{"NonsensicalterminalName---___", false},
		{"Eterm", true},
		{"gnome", true},
		{"xterms-sun", true},
	}

	for _, test := range tests {
		hc := HasColor(test.Name)
		if hc != test.HasColor {
			t.Errorf("case %q: want %t got %t", test.Name, test.HasColor, hc)
		}
	}
}
