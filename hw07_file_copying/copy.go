package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/andybalholm/crlf"
	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(from, to string, offset, limit int, withBar bool) error {
	content, bar, err := ReadFile(from, limit, offset, withBar)
	if err != nil {
		return err
	}

	wFile, err := os.OpenFile(to, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrUnsupportedFile, err)
	}
	defer wFile.Close()

	nw := crlf.NewWriter(wFile)
	if withBar {
		nw = bar.NewProxyWriter(nw)
	}

	_, err = nw.Write(content)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrUnsupportedFile, err)
	}

	if withBar {
		bar.Finish()
	}

	return nil
}

func ReadFile(from string, limit int, offset int, withBar bool) ([]byte, *pb.ProgressBar, error) {
	rFile, err := os.Open(from)
	if withBar {
		time.Sleep(1000 * time.Millisecond)
	}
	if err != nil {
		return nil, nil, ErrUnsupportedFile
	}
	defer rFile.Close()

	stat, err := rFile.Stat()
	if err != nil {
		return nil, nil, ErrUnsupportedFile
	}
	contentSize := int(stat.Size())
	if limit != 0 {
		if limit+offset >= contentSize {
			contentSize -= offset
		} else {
			contentSize = limit
		}
	}
	if contentSize <= 0 {
		return nil, nil, fmt.Errorf("content size was %v :%w", contentSize, ErrOffsetExceedsFileSize)
	}

	var bar *pb.ProgressBar
	if withBar {
		bar = pb.Full.Start(contentSize * 2)
	}

	content := make([]byte, contentSize)
	reader := crlf.NewReader(rFile)
	content = DynamicRead(contentSize, offset, reader, content, bar)
	if withBar {
		bar.SetTotal(int64(len(content) * 2))
	}

	return content, bar, err
}

func DynamicRead(limit, offset int, reader io.Reader, content []byte, bar *pb.ProgressBar) []byte {
	offsetBytes := make([]byte, offset)
	wasRead := 0
	for wasRead < limit {
		read := 0
		var err error
		if offset > 0 {
			read, err = reader.Read(offsetBytes)
			if offset >= read {
				offset -= read
				read = 0
			} else {
				bytes := offsetBytes[offset:read]
				read -= offset
				bytes = append(bytes, content[read:]...)
				content = bytes
				offset = 0
			}
		} else {
			read, err = reader.Read(content[wasRead:])
		}

		wasRead += read
		if err != nil {
			if errors.Is(err, io.EOF) {
				content = content[:wasRead]
				break
			}
			log.Panicf("failed to read: %v", err)
		}

		if bar != nil {
			bar.Add(read)
			time.Sleep(500 * time.Millisecond)
		}
	}
	return content
}
