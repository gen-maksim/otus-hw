package main

import (
	"flag"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	if limit > 0 {
		limit++
	}

	rFile, err := os.Open(from)
	if err != nil {
		return
	}
	defer rFile.Close()

	stat, _ := rFile.Stat()

	barSize := int(stat.Size())
	if limit != 0 {
		barSize = int(limit)
		if (limit - (stat.Size() - offset)) > 0 {
			barSize = int(limit - (stat.Size() - offset))
		}
	}

	bar := pb.StartNew(barSize * 2)

	content := make([]byte, barSize)
	_, err = rFile.Seek(offset, io.SeekStart)
	if err != nil {
		return
	}

	barReader := bar.NewProxyReader(rFile)
	_, err = barReader.Read(content)
	wFile, err := os.OpenFile(to, os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer wFile.Close()

	barWriter := bar.NewProxyWriter(wFile)
	_, err = barWriter.Write(content)
	if err != nil {
		return
	}
	bar.Finish()
}
