package main

import (
	"strings"
	"testing"
)

const (
	msg = "It was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness, it was the epoch of belief, it was the epoch of incredulity, " +
		"it was the season of Light, it was the season of Darkness, it was the spring of hope, it was the winter of despair, we had everything before us, we had nothing before us, " +
		"we were all going direct to Heaven, we were all going direct the other way--in short, the period was so far like the present period, that some of its noisiest authorities insisted " +
		"on its being received, for good or for evil, in the superlative degree of comparison only."
)

func Test_hashMd5(t *testing.T) {
	h := hashMd5(strings.NewReader(msg))
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

func Test_hashSha256(t *testing.T) {
	h := hashSha256(msg)
	if h != "f75e92f657ee56b86c0db88d8fb09436c7b9afbeca2df1c256b072f6eb55d766" {
		t.Errorf("incorrect hash: %s", h)
	}
}

func Test_hashSha224(t *testing.T) {
	h := hashSha224(msg)
	if h != "b7a9d63b822728eea25a335305fce43ac6103bfc5fd81db872a6b4ad" {
		t.Errorf("incorrect hash: %s", h)
	}
}

func Test_hashSha512(t *testing.T) {
	h := hashSha512(msg)
	if h != "fccea2ee8a556ef37c5f201373469f84b119ebada1cb40790e6e006224dc838fb17f676d24838707d22bcf5ed6029b3c7cf9c980625f094fa8aa1cc7206c1719" {
		t.Errorf("incorrect hash: %s", h)
	}
}

func Test_hashSha384(t *testing.T) {
	h := hashSha384(msg)
	if h != "377161ec7c3ded82d75f2c87195489f15e4901f7a7ac55ac4092db82a1d78830a28ce7d28811fd801c800cefcc44a136" {
		t.Errorf("incorrect hash: %s", h)
	}
}

func Test_hashSha512_224(t *testing.T) {
	h := hashSha512_224(msg)
	if h != "f6098184cc2de571ceee61bd252325e26a3cd4810fc1461bddef3770" {
		t.Errorf("incorrect hash: %s", h)
	}
}

func Test_hashSha512_256(t *testing.T) {
	h := hashSha512_256(msg)
	if h != "e99ae0c7e98050096288bb4d579036dcba2f01a208f6a066035672f0d512b910" {
		t.Errorf("incorrect hash: %s", h)
	}
}
