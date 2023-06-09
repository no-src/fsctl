package command

import (
	"fmt"

	"github.com/no-src/log"
)

var (
	allCommands = make(map[string]func(a Action) (Command, error))
)

// Commands the commands that contain the init, actions and clear stages
type Commands struct {
	Name    string
	Init    []Command
	Actions []Command
	Clear   []Command
}

// Exec execute the commands in order of stages
func (cs *Commands) Exec() (err error) {
	if err = cs.ExecInit(); err != nil {
		return err
	}
	if err = cs.ExecActions(); err != nil {
		return err
	}
	return cs.ExecClear()
}

// ExecInit execute the init commands
func (cs *Commands) ExecInit() (err error) {
	return cs.exec("init", cs.Init)
}

// ExecActions execute the actions commands
func (cs *Commands) ExecActions() (err error) {
	return cs.exec("actions", cs.Actions)
}

// ExecClear execute the clear commands
func (cs *Commands) ExecClear() (err error) {
	return cs.exec("clear", cs.Clear)
}

func (cs *Commands) exec(stage string, commands []Command) (err error) {
	log.Debug("[%s] start executing", stage)
	for i, c := range commands {
		step := i + 1
		if err = c.Exec(); err != nil {
			return fmt.Errorf("[%d] [%s] [%s] [failed] %w", step, stage, c.Name(), err)
		}
		log.Debug("[%d] [%s] [%s] [ok]", step, stage, c.Name())
	}
	return nil
}

func registerCommand[T Command]() {
	allCommands[(*new(T)).Name()] = func(a Action) (Command, error) {
		return parse[T](a)
	}
}
