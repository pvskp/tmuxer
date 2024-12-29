package tmux

import (
	"github.com/pvskp/tmuxer/pkg/utils"
)

type TmuxCommand struct {
	Command string
	Args    []string
}

func SendKeys(sessionName string, keys []string) (*TmuxCommand, error) {
	return &TmuxCommand{
		Command: "send-keys",
		Args:    []string{"-t", },
	}, nil
}


func (t *TmuxCommand) Execute() (string, error) {
	fullCommand := append([]string{t.Command}, t.Args...)
	return utils.ExecuteCommand("tmux", fullCommand...)
}
