package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const configDir = ".tolo"
const dbFile = "tolo.db.json"

type Alias struct {
	Name      string `json:"name"`
	Command   string `json:"command"`
	CreatedAt string `json:"created_at"`
}

type Data struct {
	Aliases []Alias `json:"aliases"`
	mu      sync.RWMutex
}

var data *Data
var dataFile string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dataFile = filepath.Join(homeDir, configDir, dbFile)
	data = &Data{}
	load()
}

func load() error {
	data.mu.Lock()
	defer data.mu.Unlock()

	if err := os.MkdirAll(filepath.Dir(dataFile), 0755); err != nil {
		return err
	}

	content, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if len(content) == 0 {
		return nil
	}

	return json.Unmarshal(content, data)
}

func saveLocked() error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, content, 0644)
}

func save() error {
	data.mu.Lock()
	defer data.mu.Unlock()
	return saveLocked()
}

func AddAlias(name, command string) error {
	data.mu.Lock()
	defer data.mu.Unlock()

	for _, a := range data.Aliases {
		if a.Name == name {
			return fmt.Errorf("alias '%s' already exists", name)
		}
	}

	data.Aliases = append(data.Aliases, Alias{
		Name:      name,
		Command:   command,
		CreatedAt: time.Now().Format(time.RFC3339),
	})

	return saveLocked()
}

func GetAlias(name string) (Alias, bool) {
	data.mu.RLock()
	defer data.mu.RUnlock()

	for _, a := range data.Aliases {
		if a.Name == name {
			return a, true
		}
	}
	return Alias{}, false
}

func UpdateAlias(name, command string) error {
	data.mu.Lock()
	defer data.mu.Unlock()

	for i, a := range data.Aliases {
		if a.Name == name {
			data.Aliases[i].Command = command
			return saveLocked()
		}
	}
	return fmt.Errorf("alias '%s' not found", name)
}

func DeleteAlias(name string) error {
	data.mu.Lock()
	defer data.mu.Unlock()

	for i, a := range data.Aliases {
		if a.Name == name {
			data.Aliases = append(data.Aliases[:i], data.Aliases[i+1:]...)
			return saveLocked()
		}
	}
	return fmt.Errorf("alias '%s' not found", name)
}

func ListAliases() []Alias {
	data.mu.RLock()
	defer data.mu.RUnlock()

	return data.Aliases
}

func ShowAlias(name string) (Alias, bool) {
	data.mu.RLock()
	defer data.mu.RUnlock()

	for _, a := range data.Aliases {
		if a.Name == name {
			return a, true
		}
	}
	return Alias{}, false
}

func SearchAliases(query string) []Alias {
	data.mu.RLock()
	defer data.mu.RUnlock()

	var results []Alias
	for _, a := range data.Aliases {
		if contains(a.Name, query) || contains(a.Command, query) {
			results = append(results, a)
		}
	}
	return results
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
