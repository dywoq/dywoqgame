package stream

import (
	"errors"
)

// Stream is a structure for containing the messages temporarily, with the maximum length.
// Its purpose is for logging, usually for games.
// The stream's messages are printed only once.
type Stream struct {
	maxLen int
	buf    []message
}

type message struct {
	bytes   []byte
	printed bool
}

var ErrFull = errors.New("stream is full")
var ErrOutOfBounds = errors.New("out of bounds")

// New returns a pointer to Stream with maxLen.
func New(maxLen int) *Stream {
	return &Stream{maxLen: maxLen, buf: make([]message, 0, maxLen)}
}

// Write writes a message in bytes.
// Returns ErrFull if the buffer is overflowed (in other words, reached the max length).
func (s *Stream) Write(p []byte) (int, error) {
	if len(s.buf) >= s.maxLen {
		return 0, ErrFull
	}
	s.buf = append(s.buf, message{p, false})
	return len(p), nil
}

// WriteByte is equal to Write([]byte(b)).
func (s *Stream) WriteByte(b byte) error {
	_, err := s.Write([]byte{b})
	return err
}

// WriteString is equal to Write([]byte(str)).
func (s *Stream) WriteString(str string) (int, error) {
	return s.Write([]byte(str))
}

// Len returns the length of the current buffer.
func (s *Stream) Len() int {
	return len(s.buf)
}

// At returns bytes and printed indicator of the message at i from the buffer.
// If i is out of bounds, the function returns nil, false and ErrOutOfBounds.
func (s *Stream) At(i int) ([]byte, bool, error) {
	if i >= len(s.buf) || i < 0 {
		return nil, false, ErrOutOfBounds
	}
	m := s.buf[i]
	return m.bytes, m.printed, nil
}

// MarkPrinted sets the printed flag for the message at index i.
// If i is out of bounds, the function returns nil, false and ErrOutOfBounds.
func (s *Stream) MarkPrinted(i int) error {
	if i >= len(s.buf) || i < 0 {
		return ErrOutOfBounds
	}
	s.buf[i].printed = true
	return nil
}
