package jaegerstart

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func StartNewSpan(ctx context.Context, operationName string) opentracing.Span {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return nil
	}

	return span.Tracer().StartSpan(operationName, opentracing.ChildOf(span.Context()))
}
