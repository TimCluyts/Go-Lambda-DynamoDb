package main

import (
	"testing"
)

func TestLambdaEvent(t *testing.T) {
	name := "Bogus"
	want := "Hello there, Bogus"
	msg, err := HandleLambdaEvent(&MyEvent{Name: name})
	if !(msg.Message == want) || err != nil {
		t.Fatalf(`Did not match strings; %s; %s; %s`, msg, err, want)
	}
}
