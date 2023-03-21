package main

import (
	"context"
	"strconv"
	"sync"

	functionsdk "github.com/numaproj/numaflow-go/pkg/function"
	"github.com/numaproj/numaflow-go/pkg/function/server"
)

var genKey = 0
var keyLock sync.RWMutex

func handle(_ context.Context, key string, d functionsdk.Datum) functionsdk.Messages {
	msg := d.Value()
	_ = d.EventTime() // Event time is available
	_ = d.Watermark() // Watermark is available
	keyLock.RLock()
	if genKey == 100 {
		keyLock.RUnlock()
		keyLock.Lock()
		genKey = 0
		keyLock.Unlock()
	} else {
		keyLock.RUnlock()
	}
	keyLock.Lock()
	genKey++
	keyLock.Unlock()
	return functionsdk.MessagesBuilder().Append(functionsdk.MessageTo(strconv.Itoa(genKey), msg))
}

func main() {
	server.New().RegisterMapper(functionsdk.MapFunc(handle)).Start(context.Background())
}
