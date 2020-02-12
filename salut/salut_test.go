package salut

import (
	"testing"
)

func TestSalut(t *testing.T) {
	cases := []struct {
		input int
		want  string
	}{
		{0, "Madame"},
		{1, "Monsieur"},
		{2, ""},
	}

	for _, c := range cases {
		out := Salut(c.input)
		if out != c.want {
			t.Errorf("Salut(%q) == %q \n", out, out)
		}
	}
}
