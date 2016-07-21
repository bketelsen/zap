package main

import (
	"github.com/bketelsen/zap/another"
	"github.com/bketelsen/zap/sub"
	"github.com/uber-go/zap"
)

func main() {

	// Log in JSON, using zap's reflection-free JSON encoder.
	// The default options will log any Info or higher logs to standard out.
	logger := zap.NewJSON()

	logger.Warn("Log from main")
	logger.Warn(
		"structured context from main",
		zap.String("application", "zap"),
	)

	sublog := logger.With(zap.String("pkg", "sub"))
	sub.Logger = sublog

	anotherlog := logger.With(zap.String("pkg", "another"))
	another.Logger = anotherlog

	sub.Sub()
	another.Another()
	logger.Info("Done.")
}
