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
		AccountID: "1234567890",
	}
	r.HandleRequest(ctx, req)
}
