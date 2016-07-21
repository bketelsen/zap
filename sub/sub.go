package sub

import "github.com/uber-go/zap"

var Logger zap.Logger

func Sub() {
	Logger.Error("something went wrong!")
}
