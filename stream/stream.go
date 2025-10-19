package stream

import (
	"errors"
	"io"
	"os"
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

// Print writes the first unwritten message to the set writer without a newline.
// It marks the message as written upon success.
// Returns true if a message was written, false otherwise.
func (w *Writer) Print() error {
	if w.s == nil || w.s.Empty() {
		return nil
	}
	for i, m := range w.s.buf {
		if !m.written {
			_, err := w.w.Write(m.bytes)
			if err != nil {
				return err
			}
			w.s.buf[i].written = true
			return nil
		}
	}
	return nil
}

// Println writes the first unwritten message to the set writer with a newline.
// It uses Print() for the main message writing.
func (w *Writer) Println() error {
	err := w.Print()
	if err != nil {
		return err
	}
	_, err = w.w.Write([]byte{'\n'})
	return err
}

// SetOutput sets out to the underlying writer.
func (w *Writer) SetOutput(out io.Writer) {
	w.w = out
}

// Print is equal to:
//
//	w := NewWriter(os.Stdout, s)
//	w.Print()
func Print(s *Stream) error {
	w := NewWriter(os.Stdout, s)
	return w.Print()
}

// Print is equal to:
//
//	w := NewWriter(os.Stdout, s)
//	w.Println()
func Println(s *Stream) error {
	w := NewWriter(os.Stdout, s)
	return w.Println()
}

// Error is equal to:
//
//	w := NewWriter(os.Stderr, s)
//	w.Print()
func Error(s *Stream) error {
	w := NewWriter(os.Stderr, s)
	return w.Print()
}

// Errorln is equal to:
//
//	w := NewWriter(os.Stderr, s)
//	w.Println()
func Errorln(s *Stream) error {
	w := NewWriter(os.Stderr, s)
	return w.Println()
}
