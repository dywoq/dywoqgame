package stream

import (
	"errors"
	"io"
)

// Stream is a structure for containing the messages temporarily, with the maximum length.
// Its purpose is for logging, usually for games.
// The stream's messages are printed only once.
type Stream struct {
	maxLen int
	buf    []*message
}

type message struct {
	bytes   []byte
	written bool
}

var ErrOutOfBounds = errors.New("stream: out of bounds")

// NewStream returns a pointer to Stream with maxLen.
func NewStream(maxLen int) *Stream {
	return &Stream{maxLen: maxLen, buf: make([]*message, 0, maxLen)}
}

// NewWriter returns a pointer to Writer.
func NewWriter(w io.Writer, s *Stream) *Writer {
	return &Writer{w, s}
}

// Write writes a message in bytes.
// If the buffer is over s.maxLen, the write methods removes the older message.
func (s *Stream) Write(p []byte) (int, error) {
	if len(s.buf) >= s.maxLen {
		s.buf = s.buf[1:]
	}
	data := make([]byte, len(p))
	copy(data, p)

	s.buf = append(s.buf, &message{data, false})
	return len(p), nil
}

// Len returns the length of the current buffer.
func (s *Stream) Len() int {
	return len(s.buf)
}

// Empty returns true if the underlying buffer is empty.
func (s *Stream) Empty() bool {
	return len(s.buf) == 0
}

// At returns bytes and written indicator of the message at i from the buffer.
// If i is out of bounds, the function returns nil, false and ErrOutOfBounds.
func (s *Stream) At(i int) ([]byte, bool, error) {
	if i >= len(s.buf) || i < 0 {
		return nil, false, ErrOutOfBounds
	}
	m := s.buf[i]
	return m.bytes, m.written, nil
}

// MarkWritten sets the written flag for the message at index i.
// If i is out of bounds, the function returns nil, false and ErrOutOfBounds.
func (s *Stream) MarkWritten(i int) error {
	if i >= len(s.buf) || i < 0 {
		return ErrOutOfBounds
	}
	s.buf[i].written = true
	return nil
}

// Clears clears the stream buffer.
func (s *Stream) Clear() {
	s.buf = s.buf[:0]
}

// Close sets nil to the underlying stream buffer.
func (s *Stream) Close() error {
	s.buf = nil
	return nil
}

// Written returns true if all messages from the buffer are written.
// If the buffer is empty, it returns false.
func (s *Stream) Written() bool {
	if s.Empty() {
		return false
	}
	for _, msg := range s.buf {
		if !msg.written {
			return false
		}
	}
	return true
}
