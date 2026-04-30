package script

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Script struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Steps       []Step `yaml:"steps"`
}

type Step struct {
	Name        string `yaml:"name"`
	Run         string `yaml:"run"`
	Register    string `yaml:"register"`
	Condition   string `yaml:"if"`
	IgnoreError bool   `yaml:"ignore_error"`
	Wait        *Wait  `yaml:"wait"`
}

type Wait struct {
	Run      string `yaml:"run"`
	Until    string `yaml:"until"`
	Timeout  string `yaml:"timeout"`
	Interval string `yaml:"interval"`
}

func Parse(data []byte) (*Script, error) {
	var s Script
	if err := yaml.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("invalid script YAML: %w", err)
	}

	if s.Name == "" {
		return nil, fmt.Errorf("script must have a 'name' field")
	}
	if len(s.Steps) == 0 {
		return nil, fmt.Errorf("script must have at least one step")
	}

	for i, step := range s.Steps {
		if step.Name == "" {
			return nil, fmt.Errorf("step %d must have a 'name' field", i+1)
		}
		if step.Run == "" && step.Wait == nil {
			return nil, fmt.Errorf("step %d (%s) must have either 'run' or 'wait'", i+1, step.Name)
		}
		if step.Wait != nil {
			if step.Wait.Run == "" {
				return nil, fmt.Errorf("step %d (%s): wait must have a 'run' field", i+1, step.Name)
			}
			if step.Wait.Until == "" {
				return nil, fmt.Errorf("step %d (%s): wait must have an 'until' field", i+1, step.Name)
			}
		}
	}

	return &s, nil
}
