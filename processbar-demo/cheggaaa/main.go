package main

import (
	"crypto/rand"
	"github.com/cheggaaa/pb/v3"
	"io"
	"io/ioutil"
)

func main() {
	var limit int64 = 1024 * 1024 * 500

	// we will copy 500 MiB from /dev/rand to /dev/null
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard

	// start new bar
	bar := pb.Full.Start64(limit)

	// create proxy reader
	barReader := bar.NewProxyReader(reader)

	// copy from proxy reader
	io.Copy(writer, barReader)

	// finish bar
	bar.Finish()
}
