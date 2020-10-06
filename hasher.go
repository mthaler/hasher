package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"crypto/md5"
)

func main() {

	const (
		defaultAlgorithm = "md5"
		usage         = "supported hash algorithms: md5"
	)


	var algorithm string

	flag.StringVar(&algorithm, "algorithm", defaultAlgorithm, usage)
	flag.StringVar(&algorithm, "a", defaultAlgorithm, usage)

	flag.Parse()

	switch algorithm {
		case "md5":
			hashMd5("The fog is getting thicker!")
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