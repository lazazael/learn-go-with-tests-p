package main

import "testing"

func Test_Hello(t *testing.T) {
	want := "Hello, world"
	got := Hello("world")

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
