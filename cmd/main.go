package main

import (
	"fmt"
	"log"

	"github.com/pvskp/tmuxer/pkg/utils"
)

func main () {
  output, err := utils.ExecuteCommand("tmux", "list-session")
  if err != nil {
    log.Fatalf("Failed to execute command: %v", err)
  }
  fmt.Println(output)
}
