package engine

import (
	"sync"

	"github.com/ExPl0iT-29/wardenGo/internal/models"
	"github.com/sirupsen/logrus"
)

type RulesEngine struct {
	Rules []models.Rule
	mu    sync.RWMutex
}

func NewRulesEngine(rules []models.Rule) *RulesEngine {
	return &RulesEngine{Rules: rules}
}

func (re *RulesEngine) UpdateRules(rules []models.Rule) {
	re.mu.Lock()
	re.Rules = rules
	re.mu.Unlock()
	logrus.Infof("Updated rules engine with %d rules", len(rules))
}

func (re *RulesEngine) Evaluate(event models.Event) {
	re.mu.RLock()
	defer re.mu.RUnlock()
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
				"rule_id":  rule.ID,
				"event_id": event.ID,
				"severity": rule.Severity,
			}).Warn("Suspicious activity detected")
			// Add alerting logic here later
		}
	}
}
