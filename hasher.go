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
		usage            = "supported hash algorithms: md5, sha1, sha224, sha256, sha384, sha512"
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
	case "sha512":
		hashSha512("The fog is getting thicker!")
	case "sha384":
		hashSha384("The fog is getting thicker!")
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

func hashSha512(msg string) string {
	h := sha512.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha384(msg string) string {
	h := sha512.New384()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha512_224(msg string) string {
	h := sha512.New512_224()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashSha512_256(msg string) string {
	h := sha512.New512_256()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}