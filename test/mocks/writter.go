package mocks

import (
	"bytes"
	"io"
)

type mockWriter struct {
	buffer *bytes.Buffer
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return m.buffer.Write(p)
}

func NewMockWritter(buffer *bytes.Buffer) io.Writer {
	return &mockWriter{
		buffer: buffer,
	}
}
