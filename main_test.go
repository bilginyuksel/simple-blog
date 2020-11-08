package main

import "testing"

func TestHello_ExpectHelloWorld(t *testing.T) {
	res := Hello()
	if res != "Hello world" {
		t.Error()
	}
}
