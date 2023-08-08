package command

import "os"

type symlink struct {
	Link        string `yaml:"link"`
	Dest        string `yaml:"dest"`
	IgnoreError bool   `yaml:"ignore-error"`
}

func (c symlink) Exec() error {
	err := os.Symlink(c.Dest, c.Link)
	if c.IgnoreError {
		return nil
	}
	return err
}

func (c symlink) Name() string {
	return "symlink"
}

func init() {
	registerCommand[symlink]()
}
