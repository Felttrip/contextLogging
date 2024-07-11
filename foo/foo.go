package foo

import (
	"context"
	logger "logging/contextLogger"

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

/*
Bar shows the usage of WithContextInLog to get back a Logrus Entry with the context values added to it.
This allows Fooer.Bar to add additional fields to the log entry without them needing to be in the context or known about by
the context aware logger.
*/
func (s *fooer) Bar(ctx context.Context) error {
	logger.WithContextInLog(ctx, s.logger).WithField("FooStuffOnlyFooKnowsAbout", "Baz").Info("We're in Bar")
	return nil
}
