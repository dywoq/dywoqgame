package stream

import (
	"io"
	"os"
)

// Writer is responsible for writing stream data
// to any structure that implements io.Writer.
type Writer struct {
	w io.Writer
	s *Stream
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
