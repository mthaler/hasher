package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	const (
		defaultAlgorithm = "md5"
		usage            = "supported hash algorithms: md5, sha1, sha224, sha256, sha384, sha512, sha512_224, sha512_256"
	)

	var algorithm string

	flag.StringVar(&algorithm, "algorithm", defaultAlgorithm, usage)
	flag.StringVar(&algorithm, "a", defaultAlgorithm, usage)

	flag.Parse()

	r := os.Stdin

	switch algorithm {
	case "md5":
		hashMd5(r)
	case "sha1":
		hashSha1(r)
	case "sha256":
		hashSha256(r)
	case "sha224":
		hashSha256(r)
	case "sha512":
		hashSha512(r)
	case "sha384":
		hashSha384(r)
	case "sha512_224":
		hashSha512_224(r)
	case "sha512_256":
		hashSha512_256(r)
	default:
		fmt.Printf("Unknown algorithm: %s\n", algorithm)
		os.Exit(1)
	}
}

func hashMd5(r io.Reader) string {
	h := md5.New()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha1(r io.Reader) string {
	h := sha1.New()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha256(r io.Reader) string {
	h := sha256.New()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha224(r io.Reader) string {
	h := sha256.New224()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha512(r io.Reader) string {
	h := sha512.New()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha384(r io.Reader) string {
	h := sha512.New384()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha512_224(r io.Reader) string {
	h := sha512.New512_224()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha512_256(r io.Reader) string {
	h := sha512.New512_256()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}
