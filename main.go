package main

import (
	"compress/gzip"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var flagTo = flag.String("to", "out.gzip", "gzip name")
	var flagFrom = flag.String("from", "", "target file name")

	flag.Parse()

	dist, err := os.Create(filepath.Join("data", *flagTo))
	if err != nil {
		log.Fatal(err)
	}
	defer dist.Close()
	gw := gzip.NewWriter(dist)
	defer gw.Close()
	src, err := os.Open(filepath.Join("data", *flagFrom))
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	if _, err := io.Copy(gw, src); err != nil {
		log.Fatal(err)
	}

}
