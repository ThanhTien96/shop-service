package api

import "log"

type NullWriter struct{}

func NewNullWriter() *NullWriter {
	return &NullWriter{}
}

func (b *NullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

var NoLogger = log.New(NewNullWriter(), "", 0)
