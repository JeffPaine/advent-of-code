package main

import "testing"

func TestHash(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			"HASH",
			52,
		},
		{
			"rn=1",
			30,
		},
		{
			"cm-",
			253,
		},
		{
			"qp=3",
			97,
		},
		{
			"cm=2",
			47,
		},
		{
			"qp-",
			14,
		},
		{
			"pc=4",
			180,
		},
		{
			"ot=9",
			9,
		},
		{
			"ab=5",
			197,
		},
		{
			"pc-",
			48,
		},
		{
			"pc=6",
			214,
		},
		{
			"ot=7",
			231,
		},
	}

	for _, test := range tests {
		if got := hash(test.s); got != test.want {
			t.Errorf("hash(%q) = %v, want: %v", test.s, got, test.want)
		}
	}
}
