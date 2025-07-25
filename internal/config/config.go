package config

import (
    "encoding/json"
    "os"
    "github.com/ExPl0iT-29/wardenGo/internal/models"
)

func LoadRules(filePath string) ([]models.Rule, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    var rules []models.Rule
    if err := json.Unmarshal(data, &rules); err != nil {
        return nil, err
    }
    return rules, nil
}