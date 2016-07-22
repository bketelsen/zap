package custom

import "github.com/bketelsen/zap/backlog"

type CustomType struct {
	Name  string
	Count int
}

func (c *CustomType) MarshalLog(kv backlog.KeyValue) error {
	kv.AddString("Name", c.Name)
	kv.AddInt("Count", c.Count)
	return nil
}
