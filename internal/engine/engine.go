package engine

import (
    "github.com/ExPl0iT-29/wardenGo/internal/models"
    "github.com/sirupsen/logrus"
)

type RulesEngine struct {
    Rules []models.Rule
}

func NewRulesEngine(rules []models.Rule) *RulesEngine {
    return &RulesEngine{Rules: rules}
}

func (re *RulesEngine) Evaluate(event models.Event) {
    for _, rule := range re.Rules {
        match := true
        for key, value := range rule.Conditions {
            if event.Details[key] != value {
                match = false
                break
            }
        }
        if match {
            logrus.WithFields(logrus.Fields{
                "rule_id": rule.ID,
                "event_id": event.ID,
                "severity": rule.Severity,
            }).Warn("Suspicious activity detected")
            // Add alerting logic (e.g., send email, push to external system)
        }
    }
}