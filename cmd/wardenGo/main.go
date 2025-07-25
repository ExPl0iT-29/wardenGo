package main

import (
    "os" // Remove blank identifier and use os for directory creation
    "github.com/sirupsen/logrus"
    "github.com/ExPl0iT-29/wardenGo/internal/config"
    "github.com/ExPl0iT-29/wardenGo/internal/engine"
    "github.com/ExPl0iT-29/wardenGo/internal/scanner"
)

func main() {
    // Load rules
    rules, err := config.LoadRules("rules/sample.json")
    if err != nil {
        logrus.Fatal("Failed to load rules: ", err)
    }

    // Initialize rules engine
    re := engine.NewRulesEngine(rules)

    // Initialize scanner
    s, err := scanner.NewScanner()
    if err != nil {
        logrus.Fatal("Failed to initialize scanner: ", err)
    }

    // Use a valid Windows path for the directory to watch
    watchDir := `C:\Temp` // Use raw string literals for Windows paths
    // Create the directory if it doesn't exist
    if err := os.MkdirAll(watchDir, 0755); err != nil {
        logrus.Fatal("Failed to create watch directory: ", err)
    }

    logrus.Infof("Watching directory: %s", watchDir)
    if err := s.Watch(watchDir); err != nil {
        logrus.Fatal("Failed to start watcher: ", err)
    }

    // Process events
    for event := range s.Events {
        re.Evaluate(event)
    }
}