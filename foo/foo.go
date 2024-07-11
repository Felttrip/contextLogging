package foo

import (
	"context"
	"logging/contextkeys"

	"github.com/sirupsen/logrus"
)

type Fooer interface {
	Bar(ctx context.Context) error
}

type fooer struct {
	logger *logrus.Logger
}

func NewFooer(l *logrus.Logger) Fooer {
	fooer := &fooer{
		logger: l,
	}
	return fooer
}

func (s *fooer) Bar(ctx context.Context) error {
	s.logger.WithFields(logrus.Fields{
		"AccountID":         ctx.Value(contextkeys.Key("AccountID")),
		"FieldNotInContext": ctx.Value(contextkeys.Key("notHere")),
	}).Info("We are in bar!")
	return nil
}
