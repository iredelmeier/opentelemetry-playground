package internal

type State struct {
	span *Span
}

func NewState() *State {
	return &State{
		span: NewSpan(),
	}
}

func (s *State) Span() *Span {
	return s.span
}

func (s *State) SetSpan(span *Span) {
	s.span = span
}
