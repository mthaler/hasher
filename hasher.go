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

	if len(flag.Args()) > 0 {
		for _, a := range flag.Args() {
			fmt.Printf("%s: ", a)
		}
	} else {
		hash(os.Stdin, algorithm)
	}
}

func hash(r io.Reader, algorithm string) string {
	switch algorithm {
	case "md5":
		return hashMd5(r)
	case "sha1":
		return hashSha1(r)
	case "sha256":
		return hashSha256(r)
	case "sha224":
		return hashSha256(r)
	case "sha512":
		return hashSha512(r)
	case "sha384":
		return hashSha384(r)
	case "sha512_224":
		return hashSha512_224(r)
	case "sha512_256":
		return hashSha512_256(r)
	default:
		fmt.Printf("Unknown algorithm: %s\n", algorithm)
		os.Exit(1)
		return ""
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
