package testx

import "testing"

func TestA(t *testing.T) {
	cases := []struct {
		a, b float64
		msg  string
	}{
		{1.0000000000000000000001, 0.999999999999999999999999999998, "hmmmm..."},
		{10000000000000.0, 10000000000000.0, "large float"},
		{1.0e48, 1.0e48, "large float 2"},
		{1.1, 1.1, "medium sized float 1"},
		{1.121212121212121212, 1.121212121212121212, "medium sized float 2"},
		{0.0000000000000000000000000000000000000000001, 0.0000000000000000000000000000000000000000001, "smaller float"},
		{4.123 + 2.3 * 4.3 / 2.3, 4.3 / 2.3 * 2.3 + 4.123, "some calc"},
	}
	for _, c := range cases {
		EqualFloat(t, c.a, c.b, c.msg)
	}
}
