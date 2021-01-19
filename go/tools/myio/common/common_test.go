package common

import "testing"

func TestSomething(t *testing.T) {

	// Set up the cases
	cases := []struct {
		in, want string
	}{
		{"rich", "hello, rich"},
		{"rich", "hello, rich"},
		{"rich", "hello, rich"},
	}

	// Call DoSomethingRandom with in and "check" for "want"
	for _, c := range cases {
		got := DoSomethingRandom(c.in)
		if got != c.want {
			t.Errorf("DoSomethingRandom(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
