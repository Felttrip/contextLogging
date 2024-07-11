package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Key string

var (
	AccountID  Key = Key("AccountID")
	MessageID  Key = Key("MessageID")
	RetryCount Key = Key("RetryCount")
	Headers    Key = Key("Headers")
)

func WithContextInFields(ctx context.Context, logger *logrus.Logger) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"AccountID":  ctx.Value(AccountID),
		"MessageID":  ctx.Value(MessageID),
		"RetryCount": ctx.Value(RetryCount),
		"Headers":    ctx.Value(Headers),
	})
}
func Info(ctx context.Context, logger *logrus.Logger, msg string) {
	logger.WithFields(logrus.Fields{
		"AccountID":  ctx.Value(AccountID),
		"MessageID":  ctx.Value(MessageID),
		"RetryCount": ctx.Value(RetryCount),
		"Headers":    ctx.Value(Headers),
	}).Info(msg)
}

func Error(ctx context.Context, logger *logrus.Logger, msg string) {
	logger.WithFields(logrus.Fields{
		"AccountID":  ctx.Value(AccountID),
		"MessageID":  ctx.Value(MessageID),
		"RetryCount": ctx.Value(RetryCount),
		"Headers":    ctx.Value(Headers),
	}).Error(msg)
}