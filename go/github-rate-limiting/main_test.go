package main

import (
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	d, err := time.ParseDuration("50ms")
	if err != nil {
		t.Fatal(err)
	}

	gh := &client{
		remaining: 1,
		reset:     time.Now().Add(d),
	}

	before := time.Now()
	gh.wait()
	after := time.Now()

	if after.Sub(before) >= d {
		t.Fail()
	}

	gh.remaining = 0
	before = time.Now()
	gh.wait()
	after = time.Now()

	if after.Sub(before) < d {
		t.Fail()
	}
}
