package tmux

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Window struct {
	Index       int
	SessionName string
	Name        string
	Panes       map[string]*Pane
}

func (w *Window) kill() error {
	cmd := &TmuxCommand{
		Command: "kill-window",
		Args: []string{
			"-t",
			fmt.Sprintf("%s:%d", w.SessionName, w.Index),
		},
	}

  if _, err := cmd.Execute(); err != nil {
    return err
  }

  return nil
}

func (w *Window) fetchPanes() {
	cmd := &TmuxCommand{
		Command: "list-panes",
		Args: []string{
			"-t",
			fmt.Sprintf("%s:%d", w.SessionName, w.Index),
			"-F",
			"#P",
		},
	}

	cmdOutput, err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

	for _, pane := range strings.Split(cmdOutput, "\n") {
		if pane != "" {
			paneIndex, err := strconv.Atoi(pane)
			if err != nil {
				log.Fatal("Failed to convert pane str to int")
			}

			w.Panes = append(w.Panes, &Pane{
				Index:       paneIndex,
				WindowName:  w.Name,
				SessionName: w.SessionName,
			})
		}
	}

}
