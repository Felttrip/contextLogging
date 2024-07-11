package requests

import (
	"context"
	"logging/contextkeys"
	"logging/foo"

	"github.com/sirupsen/logrus"
)

type RequestHandler struct {
	f      foo.Fooer
	logger *logrus.Logger
}

type Request struct {
	AccountID string
}

func NewRequestHandler(l *logrus.Logger) *RequestHandler {
	r := &RequestHandler{
		f:      foo.NewFooer(l),
		logger: l,
	}
	return r
}

func (r RequestHandler) HandleRequest(ctx context.Context, req Request) {
	aID := contextkeys.Key("AccountID")
	ctx = context.WithValue(ctx, aID, req.AccountID)
	r.logger.WithField("AccountID", ctx.Value(aID)).Info("Handling request")
	r.f.Bar(ctx)
}
