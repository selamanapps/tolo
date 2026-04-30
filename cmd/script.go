package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"tolo/pretty"
	"tolo/script"
)

func HandleScript(args []string) error {
	if len(args) == 0 {
		fmt.Println(ScriptHelp())
		return nil
	}

	subcommand := args[0]
	rest := args[1:]

	switch subcommand {
	case "run", "r":
		if len(rest) == 0 {
			pretty.Error("No script name specified")
			fmt.Println("Usage: tolo script run <name>")
			return fmt.Errorf("no script name specified")
		}
		return ScriptRun(strings.Join(rest, " "))

	case "list", "ls", "l":
		return ScriptList()

	case "show", "sh":
		if len(rest) == 0 {
			pretty.Error("No script name specified")
			fmt.Println("Usage: tolo script show <name>")
			return fmt.Errorf("no script name specified")
		}
		return ScriptShow(strings.Join(rest, " "))

	case "steps":
		if len(rest) == 0 {
			pretty.Error("No script name specified")
			fmt.Println("Usage: tolo script steps <name>")
			return fmt.Errorf("no script name specified")
		}
		return ScriptSteps(strings.Join(rest, " "))

	case "delete", "del", "rm", "d":
		if len(rest) == 0 {
			pretty.Error("No script name specified")
			fmt.Println("Usage: tolo script delete <name>")
			return fmt.Errorf("no script name specified")
		}
		return ScriptDelete(strings.Join(rest, " "))

	case "save", "s":
		if len(rest) == 0 {
			pretty.Error("No script name specified")
			fmt.Println("Usage: tolo script save <name> [-f <file>]")
			return fmt.Errorf("no script name specified")
		}
		return ScriptSave(rest)

	case "help", "h":
		fmt.Println(ScriptHelp())
		return nil

	default:
		return ScriptRun(strings.Join(args, " "))
	}
}

func ScriptRun(name string) error {
	name = strings.TrimSpace(name)
	return script.RunScript(name)
}

func ScriptList() error {
	scripts := script.List()

	if len(scripts) == 0 {
		pretty.Info("No scripts found")
		pretty.Newline()
		pretty.Dim("Create scripts in ~/.tolo/scripts/ as YAML files")
		pretty.Dim("Or use: tolo script save <name> -f <file>")
		return nil
	}

	pretty.Header("Saved Scripts")

	maxNameLen := 0
	for _, s := range scripts {
		if len(s.Name) > maxNameLen {
			maxNameLen = len(s.Name)
		}
	}

	for i, s := range scripts {
		desc := ""
		if s.Description != "" {
			desc = s.Description
		} else {
			desc = fmt.Sprintf("%d steps", s.StepCount)
		}

		fmt.Printf("  %s%d%s  %s%-*s%s  %s→%s  %s%s%s\n",
			"\033[2m", i+1, "\033[0m",
			"\033[1m", maxNameLen, s.Name, "\033[0m",
			"\033[36m", "\033[0m",
			"\033[32m", desc, "\033[0m")
	}

	pretty.Separator()
	fmt.Printf("  Total: ")
	pretty.Count(len(scripts))
	return nil
}

func ScriptShow(name string) error {
	name = strings.TrimSpace(name)

	data, err := script.LoadRaw(name)
	if err != nil {
		pretty.Error(err.Error())
		return err
	}

	pretty.Header(fmt.Sprintf("Script: %s", name))
	fmt.Printf("  %s%s%s\n", "\033[32m", string(data), "\033[0m")
	pretty.Separator()
	return nil
}

func ScriptSteps(name string) error {
	name = strings.TrimSpace(name)
	return script.RunScriptSteps(name)
}

func ScriptDelete(name string) error {
	name = strings.TrimSpace(name)

	if err := script.Delete(name); err != nil {
		pretty.Error(err.Error())
		return err
	}

	pretty.Deleted(fmt.Sprintf("Script '%s' deleted", name))
	return nil
}

func ScriptSave(args []string) error {
	var name string
	var filePath string

	name = args[0]

	for i := 1; i < len(args); i++ {
		if args[i] == "-f" || args[i] == "--file" {
			if i+1 < len(args) {
				filePath = args[i+1]
				i++
			}
		}
	}

	var content []byte
	var err error

	if filePath != "" {
		content, err = os.ReadFile(filePath)
		if err != nil {
			pretty.Error(fmt.Sprintf("Failed to read file: %v", err))
			return err
		}
	} else {
		content, err = io.ReadAll(os.Stdin)
		if err != nil {
			pretty.Error(fmt.Sprintf("Failed to read stdin: %v", err))
			return err
		}
	}

	if len(content) == 0 {
		pretty.Error("No script content provided")
		return fmt.Errorf("empty script content")
	}

	if err := script.Save(name, content); err != nil {
		pretty.Error(err.Error())
		return err
	}

	s, _ := script.Parse(content)
	pretty.Saved("Script saved successfully")
	pretty.Newline()
	pretty.Label("Name:        ")
	pretty.Alias(s.Name)
	pretty.Newline()
	if s.Description != "" {
		pretty.Label("Description: ")
		pretty.Command(s.Description)
	}
	pretty.Label("Steps:       ")
	pretty.Count(len(s.Steps))
	return nil
}

func ScriptHelp() string {
	return `tolo script - Multi-step script runner

Usage:
    tolo script <subcommand> [arguments]
    tolo sc <subcommand> [arguments]
    tolo sc <name>                    (shortcut for: tolo script run <name>)

Subcommands:
    run <name>      (r)   Execute a script
    list            (ls)  List all scripts
    show <name>     (sh)  Show script YAML content
    steps <name>          Show script steps without running
    delete <name>   (d)   Delete a script
    save <name> [-f <file>]  Save a script from file or stdin
    help            (h)   Show this help

Script YAML Format:
    name: my-script
    description: What this script does
    steps:
      - name: Check status
        run: gcloud compute instances describe my-vm --format="value(status)"
        register: status

      - name: Start if stopped
        if: "{{.status}} != RUNNING"
        run: gcloud compute instances start my-vm

      - name: Wait until ready
        wait:
          run: gcloud compute instances describe my-vm --format="value(status)"
          until: RUNNING
          timeout: 5m
          interval: 10s

      - name: Connect
        run: tolo r myserver

Step Options:
    run: <command>           Command to execute
    register: <var>          Save stdout to a variable
    if: <var> == <value>     Run only if condition is met (also !=)
    ignore_error: true       Continue on failure
    wait:                    Poll until condition is met
      run: <command>         Command to poll
      until: <value>         Expected output (trimmed)
      timeout: 5m            Max wait time (default: 5m)
      interval: 10s          Time between polls (default: 5s)

Examples:
    tolo sc cloud-start                    Run a script
    tolo script list                       List all scripts
    tolo script save deploy -f deploy.yaml Save from file
    cat script.yaml | tolo script save my  Save from stdin
    tolo script show cloud-start           View script
    tolo script delete old-script          Delete a script`
}
