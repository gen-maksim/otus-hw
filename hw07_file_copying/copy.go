package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(from, to string, offset, limit int, withBar bool) error {
	rFile, err := os.Open(from)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer rFile.Close()

	stat, err := rFile.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	contentSize := int(stat.Size())
	if limit != 0 {
		if limit+offset >= contentSize {
			contentSize -= offset
		} else {
			contentSize = limit
		}
	}
	if contentSize < 0 {
		return ErrOffsetExceedsFileSize
	}

	_, err = rFile.Seek(int64(offset), 0)
	if err != nil {
		return ErrOffsetExceedsFileSize
	}

	wFile, err := os.OpenFile(to, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0o666))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrUnsupportedFile, err)
	}
	defer wFile.Close()

	var bar *pb.ProgressBar
	nr := io.Reader(rFile)
	nw := io.Writer(wFile)
	if withBar {
		bar = pb.StartNew(contentSize)
		nr = bar.NewProxyReader(rFile)
	}

	_, err = io.CopyN(nw, nr, int64(contentSize))
	if err != nil {
		return err
	}

	if withBar {
		bar.Finish()
	}

	return nil
}
