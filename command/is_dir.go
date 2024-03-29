package command

import (
	"github.com/no-src/nsgo/fsutil"
)

type isDir struct {
	Source string `yaml:"source"`
	Expect bool   `yaml:"expect"`
}

func (c isDir) Exec() error {
	dir, err := fsutil.IsDir(c.Source)
	if err != nil {
		return err
	}
	if dir != c.Expect {
		err = newNotExpectedError(c, dir)
	}
	return err
}

func (c isDir) Name() string {
	return "is-dir"
}

func init() {
	registerCommand[isDir]()
}
