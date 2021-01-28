package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var output = flag.String("o", "a", "output file path")
var size = flag.Int64("s", 128, "output file size")

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	file, err := os.Create(*output)
	catch(err)
	writer := bufio.NewWriter(file)
	reader := io.LimitReader(rand.Reader, *size)
	bar := pb.Full.Start64(*size)
	barReader := bar.NewProxyReader(reader)
	_, err = io.Copy(writer, barReader)
	catch(err)
	bar.Finish()
	file.Close()
}
