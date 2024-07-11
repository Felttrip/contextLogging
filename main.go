package main

import (
	"context"
	"logging/requests"

	"github.com/sirupsen/logrus"
)

func main() {
	// logger is newed in main and can then be used by further packages
	logger := logrus.New()
	r := requests.NewRequestHandler(logger)

	//fake request as an example of things to log
	req := requests.Request{
		AccountID:  "AID_1234567890",
		MessageID:  "MID_8675309",
		RetryCount: 3,
		Headers: requests.Headers{
			ContentType: "application/json",
			Length:      123,
		},
	}

	ctx := context.Background()
	r.HandleRequest(ctx, req)
}
