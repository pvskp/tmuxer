package tmux

import (
	"fmt"
	"log"
	"strings"
)

type Session struct {
	Id      string
	Windows []*Window
}

func NewSession(name string) (*Session, error) {
	if name == "" {
		// let's tmux give the session name

		cmd := &TmuxCommand{
			Command: "new-session",
			Args:    []string{"-d", "-P"},
		}

		output, err := cmd.Execute()
		if err != nil {
			return nil, err
		}

		id := strings.Split(output, ":")[0]
		return &Session{
			Id: id,
		}, nil
	}

	cmd := &TmuxCommand{
		Command: "new-session",
		Args:    []string{"-d", "-P", "-s", name},
	}

	cmd.Execute()
	return &Session{
		Id: name,
	}, nil
}

func ListSessions() ([]*Session, error) {
	cmd := &TmuxCommand{
		Command: "list-sessions",
		Args:    []string{"-F", "#S"},
	}
	sessions, err := cmd.Execute()
	if err != nil {
		log.Fatalf("Couldn't fetch sessions: %v", err)
	}

	sessionsStringList := strings.Split(sessions, "\n")

	sessionsList := []*Session{}

	for _, sessionName := range sessionsStringList {
		if sessionName != "" {
			session := &Session{Id: sessionName}
			session.Windows = session.fetchWindows()
			sessionsList = append(sessionsList, session)
			fmt.Printf("Got: %s\n", sessionName)
		}
	}
	return sessionsList, nil
}

func KillSession(name string) error {
	if name == "" {
		return fmt.Errorf("KillSession needs a session name.")
	}

	cmd := &TmuxCommand{
		Command: "kill-session",
		Args:    []string{"-t", name},
	}

	cmd.Execute()

	return nil
}

func (s *Session) Kill() {
	cmd := &TmuxCommand{
		Command: "kill-session",
		Args:    []string{"-t", s.Id},
	}
	if _, err := cmd.Execute(); err != nil {
		log.Fatalf("Failed to kill session: %v", err)
	}
}

func (s *Session) fetchWindows() []*Window {
	return []*Window{}
}
