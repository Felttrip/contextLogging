package requests

import (
	"context"
	logger "logging/contextLogger"
	"logging/foo"

	"github.com/sirupsen/logrus"
)

type RequestHandler struct {
	f      foo.Fooer
	logger *logrus.Logger
}

type Request struct {
	AccountID  string
	MessageID  string
	RetryCount int
	Headers    Headers
}

type Headers struct {
	ContentType string
	Length      int
}

func NewRequestHandler(l *logrus.Logger) *RequestHandler {
	r := &RequestHandler{
		f:      foo.NewFooer(l),
		logger: l,
	}
	return r
}

func (r RequestHandler) HandleRequest(ctx context.Context, req Request) {
	ctx = context.WithValue(ctx, logger.AccountID, req.AccountID)
	ctx = context.WithValue(ctx, logger.MessageID, req.MessageID)
	ctx = context.WithValue(ctx, logger.RetryCount, req.RetryCount)
	ctx = context.WithValue(ctx, logger.Headers, req.Headers)
	logger.Info(ctx, r.logger, "Handling request")
	r.f.Bar(ctx)
}
