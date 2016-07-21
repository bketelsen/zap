package another

import "github.com/uber-go/zap"

var Logger zap.Logger

func Another() {
	Logger.Error("something went wrong!")
}
