package internal

type FinishSpan func(*Span)

func defaultFinishSpan(*Span) {}
