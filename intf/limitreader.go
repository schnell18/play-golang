package intf

import "io"

// LimitReaderStruct record limit reader's limit, current count and wrapped reader
type LimitReaderStruct struct {
	limitCount    int
	readCount     int
	wrappedReader io.Reader
}

func (l *LimitReaderStruct) Read(p []byte) (int, error) {
	if l.readCount+len(p) >= l.limitCount {
		return 0, io.EOF
	}
	n, err := l.wrappedReader.Read(p)
	if err != nil {
		return n, err
	}
	l.readCount += n
	return n, nil
}

// LimitReader wraps an io.Reader and limit total bytes to read
func LimitReader(r io.Reader, n int) io.Reader {
	reader := LimitReaderStruct{limitCount: n, wrappedReader: r}
	return &reader
}
