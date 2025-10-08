package stream

import (
	"errors"
	"os"
)

// Printer is responsible for printing messages,
// given by Stream.
type Printer struct {
	s *Stream
}

// NewPrinter returns a pointer to Printer.
func NewPrinter(s *Stream) *Printer {
	return &Printer{s}
}

// Printed checks if all messages were printed.
func (p *Printer) Printed() bool {
	if len(p.s.buf) == 0 {
		return false
	}
	for _, msg := range p.s.buf {
		if !msg.printed {
			return false
		}
	}
	return true
}

// Index prints a message at i from the stream buffer.
//
// Returns ErrOutOfBounds if is out of bounds.
//
// Returns an error "stream is nil" if the internal pointer to Stream structure.
//
// Returns nil if all messages were printed, the message at i was printed,
// or the printing is successful.
func (p *Printer) Index(i int) error {
	switch {
	case p.s == nil:
		return errors.New("stream is nil")
	case i >= p.s.Len(), i < 0:
		return ErrOutOfBounds
	case p.Printed():
		return nil
	}
	if !p.s.buf[i].printed {
		os.Stdout.Write(p.s.buf[i].bytes)
		p.s.buf[i].printed = true
	}
	return nil
}

// Print is equal to:
//
//	p.Index(p.s.Len() - 1)
func (p *Printer) Print() error {
	return p.Index(p.s.Len() - 1)
}

// Println is equal to:
//
//	p.Print()
//	os.Stdout.WriteString("\n")
func (p *Printer) Println() error {
	err := p.Print()
	if err != nil {
		return err
	}
	os.Stdout.WriteString("\n")
	return nil
}
