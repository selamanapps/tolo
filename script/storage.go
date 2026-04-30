package script

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const scriptsSubdir = "scripts"

type ScriptInfo struct {
	Name        string
	Description string
	StepCount   int
	ModifiedAt  time.Time
}

func scriptsDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".tolo", scriptsSubdir)
}

func findFile(name string) string {
	dir := scriptsDir()
	for _, ext := range []string{".yaml", ".yml"} {
		p := filepath.Join(dir, name+ext)
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}

func Load(name string) (*Script, error) {
	p := findFile(name)
	if p == "" {
		return nil, fmt.Errorf("script '%s' not found", name)
	}

	data, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("failed to read script: %w", err)
	}

	return Parse(data)
}

func LoadRaw(name string) ([]byte, error) {
	p := findFile(name)
	if p == "" {
		return nil, fmt.Errorf("script '%s' not found", name)
	}
	return os.ReadFile(p)
}

func Save(name string, content []byte) error {
	dir := scriptsDir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	s, err := Parse(content)
	if err != nil {
		return err
	}

	filename := filepath.Join(dir, s.Name+".yaml")
	return os.WriteFile(filename, content, 0644)
}

func Delete(name string) error {
	p := findFile(name)
	if p == "" {
		return fmt.Errorf("script '%s' not found", name)
	}
	return os.Remove(p)
}

func List() []ScriptInfo {
	dir := scriptsDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var scripts []ScriptInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".yaml" && ext != ".yml" {
			continue
		}

		name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		p := filepath.Join(dir, entry.Name())

		info, err := entry.Info()
		if err != nil {
			continue
		}

		data, err := os.ReadFile(p)
		if err != nil {
			continue
		}

		var s Script
		if err := yaml.Unmarshal(data, &s); err != nil {
			scripts = append(scripts, ScriptInfo{
				Name:       name,
				ModifiedAt: info.ModTime(),
			})
			continue
		}

		scripts = append(scripts, ScriptInfo{
			Name:        s.Name,
			Description: s.Description,
			StepCount:   len(s.Steps),
			ModifiedAt:  info.ModTime(),
		})
	}

	return scripts
}

func Exists(name string) bool {
	return findFile(name) != ""
}
