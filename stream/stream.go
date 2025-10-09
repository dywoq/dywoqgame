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

// Writer is responsible for writing stream data
// to any structure that implements io.Writer.
type Writer struct {
	w io.Writer
	s *Stream
}

type message struct {
	bytes   []byte
	written bool
}

var ErrFull = errors.New("stream is full")
var ErrOutOfBounds = errors.New("out of bounds")

// NewStream returns a pointer to Stream with maxLen.
func NewStream(maxLen int) *Stream {
	return &Stream{maxLen: maxLen, buf: make([]*message, 0, maxLen)}
}

// NewWriter returns a pointer to Writer.
func NewWriter(w io.Writer, s *Stream) *Writer {
	return &Writer{w, s}
}

// Write writes a message in bytes.
// Returns ErrFull if the buffer is overflowed (in other words, reached the max length).
func (s *Stream) Write(p []byte) (int, error) {
	if len(s.buf) >= s.maxLen {
		return 0, ErrFull
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

// Println writes the message to the set writer without newline.
// only if it wasn't written already.
//
// Does nothing if the buffer is empty.
func (w *Writer) Print() error {
	if w.s == nil || len(w.s.buf) == 0 {
		return nil
	}
	m := w.s.buf[len(w.s.buf)-1]
	if m.written {
		return nil
	}
	_, err := w.w.Write(m.bytes)
	m.written = true
	return err
}

// Println writes the message to the set writer with newline.
// only if it wasn't written already.
//
// Does nothing if the buffer is empty.
func (w *Writer) Println() error {
	err := w.Print()
	if err != nil {
		return err
	}
	_, err = w.w.Write([]byte(string('\n')))
	return err
}

// SetOutput sets out to the underlying writer.
func (w *Writer) SetOutput(out io.Writer) {
	w.w = out
}
