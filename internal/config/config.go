package config

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/ExPl0iT-29/wardenGo/internal/models"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

type RuleManager struct {
	Rules    []models.Rule
	mu       sync.RWMutex
	FilePath string
}

func NewRuleManager(filePath string) (*RuleManager, error) {
	rm := &RuleManager{FilePath: filePath}
	if err := rm.LoadRules(); err != nil {
		return nil, err
	}
	return rm, nil
}

func (rm *RuleManager) LoadRules() error {
	data, err := os.ReadFile(rm.FilePath)
	if err != nil {
		return err
	}
	var rules []models.Rule
	if err := json.Unmarshal(data, &rules); err != nil {
		return err
	}
	rm.mu.Lock()
	rm.Rules = rules
	rm.mu.Unlock()
	logrus.Infof("Loaded %d rules from %s", len(rules), rm.FilePath)
	return nil
}

func (rm *RuleManager) WatchRules(updateChan chan<- []models.Rule) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					if err := rm.LoadRules(); err != nil {
						logrus.Error("Failed to reload rules: ", err)
						continue
					}
					rm.mu.RLock()
					updateChan <- rm.Rules
					rm.mu.RUnlock()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logrus.Error("Rules watcher error: ", err)
			}
		}
	}()
	err = watcher.Add(rm.FilePath)
	if err != nil {
		watcher.Close()
		return err
	}
	return nil
}
