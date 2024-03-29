package command

import (
	"github.com/no-src/nsgo/fsutil"
)

type isExist struct {
	Source string `yaml:"source"`
	Expect bool   `yaml:"expect"`
}

func (c isExist) Exec() error {
	exist, err := fsutil.FileExist(c.Source)
	if err == nil && exist != c.Expect {
		err = newNotExpectedError(c, exist)
	}
	return err
}

func (c isExist) Name() string {
	return "is-exist"
}

func init() {
	registerCommand[isExist]()
}
