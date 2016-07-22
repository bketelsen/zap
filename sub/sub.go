package sub

import (
	"github.com/bketelsen/zap/backlog"
)

var Logger = backlog.NewJSON()

func Sub() {
	Logger.Error("something went wrong!")
	Logger.Warn("details", backlog.String("caller", "brian"))
}
