package main

import (
	"context"
	"strconv"

	functionsdk "github.com/numaproj/numaflow-go/pkg/function"
	"github.com/numaproj/numaflow-go/pkg/function/server"
)

var genKey = 0

func handle(_ context.Context, key string, d functionsdk.Datum) functionsdk.Messages {
	msg := d.Value()
	_ = d.EventTime() // Event time is available
	_ = d.Watermark() // Watermark is available
	if genKey == 10 {
		genKey = 0
	}
	genKey++
	return functionsdk.MessagesBuilder().Append(functionsdk.MessageTo(strconv.Itoa(genKey), msg))
}

func main() {
	server.New().RegisterMapper(functionsdk.MapFunc(handle)).Start(context.Background())
}
