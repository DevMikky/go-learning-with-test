package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello("Mikky")
	want := "Hello, Mikky"

	if got != want {

		t.Errorf("got %q; want %q", got, want)
	}
}
