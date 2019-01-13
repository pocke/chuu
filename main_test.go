package main

import (
	"bytes"
	"testing"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		src      string
		expected string
	}{
		{"foo", "foo"},
		{"foo(注:aaa)", "foo[^1]\n[^1]: aaa\n"},
		{"foo(注:aaa)bar", "foo[^1]bar\n[^1]: aaa\n"},
		{"foo(注:aaa)bar(注:bbb)", "foo[^1]bar[^2]\n[^1]: aaa\n[^2]: bbb\n"},
	}

	for _, v := range cases {
		r := bytes.NewReader([]byte(v.src))
		w := bytes.NewBuffer([]byte{})
		err := Convert(r, w)
		if err != nil {
			t.Fatal(err)
		}
		if w.String() != v.expected {
			t.Errorf("expected %s, but got %s", v.expected, w.String())
		}
	}
}
