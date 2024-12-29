package main

import (
	"fmt"

	"github.com/pvskp/tmuxer/pkg/tmux"
)

func main() {
	sessions, _ := tmux.ListSessions()
}
