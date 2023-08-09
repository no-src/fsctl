package command

import (
	"os"
)

type mkdir struct {
	Source   string      `yaml:"source"`
	Perm     os.FileMode `yaml:"perm"`
	MustPerm os.FileMode `yaml:"must-perm"`
}

func (c mkdir) Exec() error {
	perm := c.Perm
	if perm == 0 {
		perm = defaultDirPerm
	}
	err := os.MkdirAll(c.Source, perm)
	if err == nil && c.MustPerm > 0 {
		return os.Chmod(c.Source, c.MustPerm)
	}
	return err
}

func (c mkdir) Name() string {
	return "mkdir"
}

func init() {
	registerCommand[mkdir]()
}
