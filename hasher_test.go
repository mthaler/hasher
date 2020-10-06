package main

import (
	"testing"
)

const (
	msg = "It was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness, it was the epoch of belief, it was the epoch of incredulity, " +
		"it was the season of Light, it was the season of Darkness, it was the spring of hope, it was the winter of despair, we had everything before us, we had nothing before us, " +
		"we were all going direct to Heaven, we were all going direct the other way--in short, the period was so far like the present period, that some of its noisiest authorities insisted " +
		"on its being received, for good or for evil, in the superlative degree of comparison only."
)

func Test_hashMd5(t *testing.T) {
	h := hashMd5(msg)
	if h != "66bbc653e921587b05cd55fd3344de4d" {
		t.Errorf("incorrect hash: %s", h)
	}
}

func Test_hashSha1(t *testing.T) {
	h := hashSha1(msg)
	if h != "a123c19ce8b485d0e42594f57fbbd82b71bbe498" {
		t.Errorf("incorrect hash: %s", h)
	}
}
