package another

import "github.com/bketelsen/zap/backlog"

var Logger = backlog.NewJSON()

func Another() {
	Logger.Error("something went wrong!")
}
