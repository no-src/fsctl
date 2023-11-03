package command

import (
	"github.com/no-src/nsgo/fsutil"
)

type isSymlink struct {
	Link        string `yaml:"link"`
	Expect      bool   `yaml:"expect"`
	IgnoreError bool   `yaml:"ignore-error"`
}

func (c isSymlink) Exec() error {
	actual, err := fsutil.IsSymlink(c.Link)
	if err != nil {
		if c.IgnoreError {
			return nil
		}
		return err
	}
	if c.Expect != actual {
		return newNotExpectedError(c, actual)
	}
	return nil
}

func (c isSymlink) Name() string {
	return "is-symlink"
}

func init() {
	registerCommand[isSymlink]()
}
