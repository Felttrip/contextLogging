package main

import (
	"context"
	"logging/requests"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	logger := logrus.New()
	r := requests.NewRequestHandler(logger)
	req := requests.Request{
		AccountID:  "AID_1234567890",
		MessageID:  "MID_8675309",
		RetryCount: 3,
		Headers: requests.Headers{
			ContentType: "application/json",
			Length:      123,
		},
	}
	r.HandleRequest(ctx, req)
}
