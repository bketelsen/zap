package main

import (
	"github.com/bketelsen/zap/another"
	"github.com/bketelsen/zap/backlog"
	"github.com/bketelsen/zap/custom"
	"github.com/bketelsen/zap/sub"
)

func main() {

	logger := backlog.NewJSON(backlog.AddCaller())

	logger.Warn("Log from main")

	logger.WithString("mykey", "myvalue").Warn("message is clear")

	sublog := logger.WithString("pkg", "sub")
	sub.Logger = sublog

	anotherlog := logger.WithString("pkg", "another")
	another.Logger = anotherlog

	sub.Sub()
	another.Another()

	ct := &custom.CustomType{
		Name:  "Server",
		Count: 5,
	}
	sublog.WithType("custom", ct).Info("Broken thing")
	logger.Info("Done.")
}
