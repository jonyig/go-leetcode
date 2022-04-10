package receiver_interface

type S struct {
	Data string
}

func (s S) Read() string {
	return s.Data
}

func (s *S) Write(str string) {
	s.Data = str
}
