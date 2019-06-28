package internal

import "context"

type FinishSpan func(context.Context, *Span)

func defaultFinishSpan(context.Context, *Span) {}
