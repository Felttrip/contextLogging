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

func (s *fooer) Bar(ctx context.Context) error {
	logger.WithContextInFields(ctx, s.logger).WithField("FooStuffOnlyFooKnowsAbout", "Baz").Info("We're in Bar")
	return nil
}
