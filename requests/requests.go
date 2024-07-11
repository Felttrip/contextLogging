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

/*
NewRequestHandler creates a new RequestHandler with a new Fooer that shares the same logger that was provided from main.
*/
func NewRequestHandler(l *logrus.Logger) *RequestHandler {
	r := &RequestHandler{
		f:      foo.NewFooer(l),
		logger: l,
	}
	return r
}

/*
HandleRequest show usage of PackContext to add a bunch if items to the context at once.
It then shows standard usage of logger.Info to log the context values with a message before calling Fooer.Bar.
*/
func (r RequestHandler) HandleRequest(ctx context.Context, req Request) {
	ctx = logger.PackContext(ctx, logger.Fields{
		logger.AccountID:  req.AccountID,
		logger.MessageID:  req.MessageID,
		logger.RetryCount: req.RetryCount,
		logger.Headers:    req.Headers,
	})
	logger.Info(ctx, r.logger, "Handling request")
	r.f.Bar(ctx)
}
