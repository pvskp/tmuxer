package main

import (
	"fmt"
	"log"

	"github.com/pvskp/tmuxer/pkg/tmux"
)

func main () {
  cmd, err := tmux.KillSession("tmuxinator")
  if err != nil {
    log.Fatal(err)
  }
  output, err := cmd.Execute()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(output)
}
