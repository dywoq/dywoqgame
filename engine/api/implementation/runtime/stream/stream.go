package stream

type Stream struct {
	data []byte
	max  int
}

func New(max int) *Stream {
	s := &Stream{make([]byte, max), max}
	return s
}

func (s *Stream) Write(p []byte) (int, error) {
	if s.data == nil {
		return 0, ErrDataNil
	}
	if len(s.data) >= s.max {
		return 0, ErrDataFull
	}
	if len(s.data)+len(p) > s.max {
		p = p[:s.max-len(s.data)]
	}
	s.data = append(s.data, p...)
	return len(p), nil
}

func (s *Stream) Flush() ([]byte, error) {
	if s.data == nil {
		return nil, ErrDataNil
	}
	flushed := s.data
	s.data = s.data[:0]
	return flushed, nil
}

func (s *Stream) Clean() error {
	if s.data == nil {
		return ErrDataNil
	}
	s.data = s.data[:0]
	return nil
}

func (s *Stream) Get() ([]byte, error) {
	if s.data == nil {
		return nil, ErrDataNil
	}
	res := make([]byte, len(s.data))
	copy(res, s.data)
	return res, nil
}

func (s *Stream) Capacity() (int, error) {
	if s.data == nil {
		return 0, ErrDataNil
	}
	return cap(s.data), nil
}

func (s *Stream) Length() (int, error) {
	if s.data == nil {
		return 0, ErrDataNil
	}
	return len(s.data), nil
}
