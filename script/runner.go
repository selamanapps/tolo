package script

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
	"tolo/pretty"
)

type Runner struct {
	vars map[string]string
}

func NewRunner() *Runner {
	return &Runner{
		vars: make(map[string]string),
	}
}

func (r *Runner) Run(s *Script) error {
	total := len(s.Steps)
	pretty.Header(fmt.Sprintf("Running Script: %s", s.Name))
	if s.Description != "" {
		pretty.Dim("  " + s.Description)
		pretty.Newline()
	}

	for i, step := range s.Steps {
		num := i + 1
		pretty.Label(fmt.Sprintf("  Step %d/%d: ", num, total))
		pretty.Alias(step.Name)
		pretty.Newline()

		if step.Condition != "" {
			ok, err := evaluateCondition(step.Condition, r.vars)
			if err != nil {
				return fmt.Errorf("step %d (%s): condition error: %w", num, step.Name, err)
			}
			if !ok {
				pretty.Dim("    Skipped (condition not met)")
				if num < total {
					pretty.Newline()
				}
				continue
			}
		}

		var stepErr error
		if step.Wait != nil {
			stepErr = r.runWait(step.Wait, num, total)
		} else if step.Run != "" {
			stepErr = r.runCommand(step.Run, step.Register)
		}

		if stepErr != nil {
			if !step.IgnoreError {
				return fmt.Errorf("step %d (%s): %w", num, step.Name, stepErr)
			}
			pretty.Dim(fmt.Sprintf("    Warning: %v (ignored)", stepErr))
		}

		if num < total {
			pretty.Newline()
		}
	}

	pretty.Newline()
	pretty.Saved("Script completed successfully")
	return nil
}

func (r *Runner) runCommand(command string, register string) error {
	cmd := substituteVars(command, r.vars)
	pretty.Dim(fmt.Sprintf("    Running: %s", cmd))

	output, err := execute(cmd)
	if err != nil {
		return err
	}

	output = strings.TrimSpace(output)
	if output != "" {
		fmt.Printf("    %sOutput:%s %s\n", "\033[2m", "\033[0m", output)
	}

	if register != "" {
		r.vars[register] = output
		pretty.Dim(fmt.Sprintf("    Saved: %s = %s", register, output))
	}

	return nil
}

func (r *Runner) runWait(w *Wait, stepNum, total int) error {
	timeout := 5 * time.Minute
	interval := 5 * time.Second

	if w.Timeout != "" {
		d, err := time.ParseDuration(w.Timeout)
		if err != nil {
			return fmt.Errorf("invalid timeout '%s': %w", w.Timeout, err)
		}
		timeout = d
	}
	if w.Interval != "" {
		d, err := time.ParseDuration(w.Interval)
		if err != nil {
			return fmt.Errorf("invalid interval '%s': %w", w.Interval, err)
		}
		interval = d
	}

	cmd := substituteVars(w.Run, r.vars)
	until := substituteVars(w.Until, r.vars)

	pretty.Dim(fmt.Sprintf("    Waiting (timeout: %s, interval: %s)...", timeout, interval))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	start := time.Now()
	for {
		output, err := execute(cmd)
		if err != nil {
			output = ""
		}
		output = strings.TrimSpace(output)

		if output == until {
			elapsed := time.Since(start).Truncate(time.Second)
			pretty.Dim(fmt.Sprintf("    Condition met after %s: %s", elapsed, output))
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout after %s (last output: '%s', expected: '%s')",
				timeout, output, until)
		case <-time.After(interval):
			continue
		}
	}
}

func execute(command string) (string, error) {
	parts := parseCommand(command)
	if len(parts) == 0 {
		return "", fmt.Errorf("empty command")
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%w: %s", err, strings.TrimSpace(stderr.String()))
	}

	return stdout.String(), nil
}

func evaluateCondition(cond string, vars map[string]string) (bool, error) {
	cond = strings.TrimSpace(cond)

	if strings.Contains(cond, "!=") {
		parts := strings.SplitN(cond, "!=", 2)
		left := strings.TrimSpace(substituteVars(parts[0], vars))
		right := strings.TrimSpace(parts[1])
		return left != right, nil
	}

	if strings.Contains(cond, "==") {
		parts := strings.SplitN(cond, "==", 2)
		left := strings.TrimSpace(substituteVars(parts[0], vars))
		right := strings.TrimSpace(parts[1])
		return left == right, nil
	}

	return false, fmt.Errorf("invalid condition '%s' (use == or !=)", cond)
}

func substituteVars(s string, vars map[string]string) string {
	for k, v := range vars {
		s = strings.ReplaceAll(s, "{{."+k+"}}", v)
	}
	return s
}

func parseCommand(command string) []string {
	var parts []string
	var current strings.Builder
	var inSingleQuote, inDoubleQuote bool

	for _, r := range command {
		switch {
		case r == '\'' && !inDoubleQuote:
			inSingleQuote = !inSingleQuote
		case r == '"' && !inSingleQuote:
			inDoubleQuote = !inDoubleQuote
		case r == ' ' && !inSingleQuote && !inDoubleQuote:
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}

func RunScript(name string) error {
	s, err := Load(name)
	if err != nil {
		return err
	}

	runner := NewRunner()
	return runner.Run(s)
}

func RunScriptSteps(name string) error {
	s, err := Load(name)
	if err != nil {
		return err
	}

	pretty.Header(fmt.Sprintf("Script: %s", s.Name))
	if s.Description != "" {
		pretty.Dim("  " + s.Description)
		pretty.Newline()
	}

	for i, step := range s.Steps {
		fmt.Printf("  %s%d.%s %s", "\033[2m", i+1, "\033[0m", step.Name)
		if step.Condition != "" {
			fmt.Printf(" %s(if: %s)%s", "\033[33m", step.Condition, "\033[0m")
		}
		if step.Wait != nil {
			fmt.Printf(" %s[wait]%s", "\033[36m", "\033[0m")
		}
		fmt.Println()
	}

	pretty.Separator()
	fmt.Printf("  Total: ")
	pretty.Count(len(s.Steps))
	return nil
}
