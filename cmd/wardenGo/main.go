package main

import (
	"flag"
	"os"
	"strings"

	"github.com/ExPl0iT-29/wardenGo/internal/config"
	"github.com/ExPl0iT-29/wardenGo/internal/engine"
	"github.com/ExPl0iT-29/wardenGo/internal/models"
	"github.com/ExPl0iT-29/wardenGo/internal/scanner"
	"github.com/sirupsen/logrus"
)

func main() {
	// Define command-line flags
	watchDirs := flag.String("watch-dirs", os.TempDir(), "Comma-separated list of directories to monitor")
	rulesFile := flag.String("rules-file", "rules/sample.json", "Path to the rules file")
	flag.Parse()

	// Initialize rule manager
	ruleManager, err := config.NewRuleManager(*rulesFile)
	if err != nil {
		logrus.Fatal("Failed to initialize rule manager: ", err)
	}

	// Initialize rules engine
	re := engine.NewRulesEngine(ruleManager.Rules)

	// Start watching rules file for changes
	ruleUpdateChan := make(chan []models.Rule)
	go func() {
		for rules := range ruleUpdateChan {
			re.UpdateRules(rules)
		}
	}()
	if err := ruleManager.WatchRules(ruleUpdateChan); err != nil {
		logrus.Fatal("Failed to watch rules file: ", err)
	}

	// Initialize scanner
	s, err := scanner.NewScanner()
	if err != nil {
		logrus.Fatal("Failed to initialize scanner: ", err)
	}

	// Split watch directories and start watching
	directories := strings.Split(*watchDirs, ",")
	for _, dir := range directories {
		dir = strings.TrimSpace(dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			logrus.Fatalf("Failed to create watch directory %s: %v", dir, err)
		}
		logrus.Infof("Watching directory: %s", dir)
		if err := s.Watch(dir); err != nil {
			logrus.Fatalf("Failed to start watcher for %s: %v", dir, err)
		}
	}

	// Process events
	for event := range s.Events {
		re.Evaluate(event)
	}
}
