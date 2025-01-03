package tmux

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Session struct {
	Id      string
	Windows map[string]*Window
}

func NewSession(name string) (*Session, error) {
	if name == "" { // let's tmux give the session name
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
			session.fetchWindows()
			sessionsList = append(sessionsList, session)
		}
	}
	return sessionsList, nil
}

func GetSession(name string) (*Session, error) {
	sessions, err := ListSessions()

	if err != nil {
		log.Fatal(err)
	}

	var session *Session
	for _, v := range sessions {
		if v.Id == name {
			session = v
			break
		}
	}

	if session == nil {
		return nil, fmt.Errorf("The given session does not exists")
	}

	session.fetchWindows()

	for _, window := range session.Windows {
		window.fetchPanes()
	}

	return session, nil
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

func (s *Session) RemoveWindow(windowName string) error {
  if w, exists := s.Windows[windowName]; exists {
    if err := w.kill(); err != nil {
      return err
    }
    delete(s.Windows, windowName)
    return nil
  }
  return fmt.Errorf("The given window does not exists on this session")
}

func (s *Session) fetchWindows() {
	cmd := &TmuxCommand{
		Command: "list-windows",
		Args:    []string{"-t", s.Id, "-F", "#I:#W"},
	}

	cmdOutput, err := cmd.Execute()

	if err != nil {
		log.Fatalf("Failed to list windows: %v\n", err)
	}

	windowStringList := strings.Split(cmdOutput, "\n")

	windowsList := []*Window{}

	for _, v := range windowStringList {
		if v != "" {
			splitOutput := strings.Split(v, ":")
			index, err := strconv.Atoi(splitOutput[0])
			if err != nil {
				log.Fatal(err)
			}
			window := &Window{
				Index:       index,
				SessionName: s.Id,
				Name:        splitOutput[1],
				Panes:       []*Pane{},
			}

			windowsList = append(windowsList, window)
		}
	}

	s.Windows = windowsList
}
