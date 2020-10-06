package main

import (
	"flag"
	"fmt"
	"os"
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
		default:
			fmt.Printf("Unknown algorithm: %s\n", algorithm)
			os.Exit(1)
	}
}
