package tmux

import (
	"fmt"

	"github.com/pvskp/tmuxer/pkg/utils"
)

type TmuxCommand struct {
	Command string
	Args    []string
}

func NewSession(name string) (*TmuxCommand, error) {
	if name == "" {
		return &TmuxCommand{
			Command: "new-session",
			Args:    []string{"-d"},
		}, nil
	}

	return &TmuxCommand{
		Command: "new-session",
		Args:    []string{"-d", "-s", name},
	}, nil
}

func ListSessions() (*TmuxCommand, error) {
	return &TmuxCommand{
		Command: "list-sessions",
		Args:    []string{"-F", "#S"},
	}, nil
}

func KillSession(name string) (*TmuxCommand, error) {
  if name == "" {
    return nil, fmt.Errorf("KillSession needs a session name.")
  }

	return &TmuxCommand{
		Command: "kill-session",
		Args:    []string{"-t", name},
	}, nil
}

func (t *TmuxCommand) Execute() (string, error) {
	fullCommand := append([]string{t.Command}, t.Args...)
	return utils.ExecuteCommand("tmux", fullCommand...)
}
