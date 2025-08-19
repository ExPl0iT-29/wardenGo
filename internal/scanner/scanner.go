package scanner

import (
	"fmt"
	"os"
	"time"

	"github.com/ExPl0iT-29/wardenGo/internal/models"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

type Scanner struct {
	Watcher *fsnotify.Watcher
	Events  chan models.Event
}

func NewScanner() (*Scanner, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Scanner{
		Watcher: watcher,
		Events:  make(chan models.Event),
	}, nil
}

func (s *Scanner) Watch(path string) error {
	// Check if the path exists and is a directory
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", path)
	}
	if err != nil {
		return fmt.Errorf("failed to stat path: %v", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", path)
	}

	err = s.Watcher.Add(path)
	if err != nil {
		return fmt.Errorf("failed to watch path: %v", err)
	}

	// Start the event loop only once (in NewScanner or main)
	// This avoids creating multiple goroutines for each Watch call
	if len(s.Watcher.WatchList()) == 1 {
		go func() {
			for {
				select {
				case event, ok := <-s.Watcher.Events:
					if !ok {
						return
					}
					logrus.Info("Detected event: ", event)
					s.Events <- models.Event{
						ID:        generateID(),
						Timestamp: time.Now(),
						Type:      "file_access",
						Details: map[string]interface{}{
							"file_path": event.Name,
							"action":    event.Op.String(),
						},
					}
				case err, ok := <-s.Watcher.Errors:
					if !ok {
						return
					}
					logrus.Error("Error: ", err)
				}
			}
		}()
	}
	return nil
}

func (s *Scanner) Close() error {
	return s.Watcher.Close()
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
