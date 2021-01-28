package main

import (
	"crypto/rand"
	"flag"
	"io/ioutil"
)

var output = flag.String("o", "a", "output file path")
var size = flag.Uint64("s", 128, "output file size")

func main() {
	flag.Parse()
	data := make([]byte, *size)
	rand.Read(data)
	ioutil.WriteFile(*output, data, 0644)
}
