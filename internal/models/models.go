package models

import (
    "time" // Add this import
)

// Event represents a system event (e.g., file access, network activity)
type Event struct {
    ID        string    `json:"id"`
    Timestamp time.Time `json:"timestamp"`
    Type      string    `json:"type"` // e.g., "file_access", "network"
    Details   map[string]interface{} `json:"details"`
}

// Rule defines a detection rule
type Rule struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Conditions  map[string]string `json:"conditions"` // e.g., {"file_path": "/etc/passwd", "action": "write"}
    Severity    string            `json:"severity"`   // e.g., "low", "medium", "high"
    Action      string            `json:"action"`     // e.g., "alert", "log"
}