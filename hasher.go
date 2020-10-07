package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	const (
		defaultAlgorithm = "md5"
		usage            = "supported hash algorithms: md5, sha1, sha224, sha256"
	)

	var algorithm string

	flag.StringVar(&algorithm, "algorithm", defaultAlgorithm, usage)
	flag.StringVar(&algorithm, "a", defaultAlgorithm, usage)

	flag.Parse()

	switch algorithm {
	case "md5":
		hashMd5("The fog is getting thicker!")
	case "sha1":
		hashSha1("The fog is getting thicker!")
	case "sha256":
		hashSha256("The fog is getting thicker!")
	case "sha224":
		hashSha256("The fog is getting thicker!")
	default:
		fmt.Printf("Unknown algorithm: %s\n", algorithm)
		os.Exit(1)
	}
}

func hashMd5(msg string) string {
	h := md5.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha1(msg string) string {
	h := sha1.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha256(msg string) string {
	h := sha256.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha224(msg string) string {
	h := sha256.New224()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}