package main

import "testing"

func TestPHello(t *testing.T) {
	got := PHello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q , want %q", got, want)
	}
}
