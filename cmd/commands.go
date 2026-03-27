package cmd

import (
	"fmt"
	"strings"
	"tolo/executor"
	"tolo/pretty"
	"tolo/storage"
)

func Save(args string) error {
	parts := strings.SplitN(args, ":", 2)
	if len(parts) != 2 {
		pretty.Error("Invalid format")
		fmt.Println("Usage: tolo save alias:command")
		return fmt.Errorf("invalid format")
	}

	alias := strings.TrimSpace(parts[0])
	command := strings.TrimSpace(parts[1])

	if alias == "" || command == "" {
		pretty.Error("Alias and command cannot be empty")
		return fmt.Errorf("empty alias or command")
	}

	if err := storage.AddAlias(alias, command); err != nil {
		pretty.Error(err.Error())
		return err
	}

	pretty.Saved(fmt.Sprintf("Alias saved successfully"))
	pretty.Newline()
	pretty.Label("Alias:   ")
	pretty.Alias(alias)
	pretty.Newline()
	pretty.Label("Command: ")
	pretty.Command(command)
	return nil
}

func Run(args string) error {
	alias := strings.TrimSpace(args)
	if alias == "" {
		pretty.Error("No alias specified")
		fmt.Println("Usage: tolo run alias")
		return fmt.Errorf("no alias specified")
	}

	aliasData, found := storage.GetAlias(alias)
	if !found {
		pretty.Error(fmt.Sprintf("Alias '%s' not found", alias))
		return fmt.Errorf("alias not found")
	}

	pretty.Separator()
	pretty.Label("Alias:   ")
	pretty.Alias(alias)
	pretty.Newline()
	pretty.Label("Command: ")
	pretty.Command(aliasData.Command)
	pretty.Separator()
	pretty.Newline()
	pretty.Running("Executing command...")
	pretty.Newline()

	return executor.Execute(aliasData.Command)
}

func Update(args string) error {
	parts := strings.SplitN(args, ":", 2)
	if len(parts) != 2 {
		pretty.Error("Invalid format")
		fmt.Println("Usage: tolo update alias:new_command")
		return fmt.Errorf("invalid format")
	}

	alias := strings.TrimSpace(parts[0])
	command := strings.TrimSpace(parts[1])

	if alias == "" || command == "" {
		pretty.Error("Alias and command cannot be empty")
		return fmt.Errorf("empty alias or command")
	}

	if err := storage.UpdateAlias(alias, command); err != nil {
		pretty.Error(err.Error())
		return err
	}

	pretty.Updated("Alias updated successfully")
	pretty.Newline()
	pretty.Label("Alias:   ")
	pretty.Alias(alias)
	pretty.Newline()
	pretty.Label("Command: ")
	pretty.Command(command)
	return nil
}

func Delete(args string) error {
	alias := strings.TrimSpace(args)
	if alias == "" {
		pretty.Error("No alias specified")
		fmt.Println("Usage: tolo delete alias")
		return fmt.Errorf("no alias specified")
	}

	if err := storage.DeleteAlias(alias); err != nil {
		pretty.Error(err.Error())
		return err
	}

	pretty.Deleted(fmt.Sprintf("Alias '%s' deleted", alias))
	return nil
}

func List(args string) error {
	aliases := storage.ListAliases()

	if len(aliases) == 0 {
		pretty.Info("No aliases found")
		pretty.Newline()
		pretty.Dim("Use 'tolo save alias:command' to add a new alias")
		return nil
	}

	pretty.Header("Saved Aliases")

	maxNameLen := 0
	maxCmdLen := 0
	for _, a := range aliases {
		if len(a.Name) > maxNameLen {
			maxNameLen = len(a.Name)
		}
		if len(a.Command) > maxCmdLen {
			maxCmdLen = len(a.Command)
		}
	}

	for i, a := range aliases {
		fmt.Printf("  %s%d%s  %s%-*s%s  %s→%s  %s%s%s\n",
			"\033[2m", i+1, "\033[0m",
			"\033[1m", maxNameLen, a.Name, "\033[0m",
			"\033[36m", "\033[0m",
			"\033[32m", a.Command, "\033[0m")
	}

	pretty.Separator()
	fmt.Printf("  Total: ")
	pretty.Count(len(aliases))
	return nil
}

func Show(args string) error {
	alias := strings.TrimSpace(args)
	if alias == "" {
		pretty.Error("No alias specified")
		fmt.Println("Usage: tolo show alias")
		return fmt.Errorf("no alias specified")
	}

	aliasData, found := storage.ShowAlias(alias)
	if !found {
		pretty.Error(fmt.Sprintf("Alias '%s' not found", alias))
		return fmt.Errorf("alias not found")
	}

	pretty.Header(fmt.Sprintf("Alias Details: %s", alias))
	fmt.Printf("  ")
	pretty.Label("Name:    ")
	pretty.Alias(aliasData.Name)
	pretty.Newline()
	fmt.Printf("  ")
	pretty.Label("Command: ")
	pretty.Command(aliasData.Command)
	pretty.Newline()
	fmt.Printf("  ")
	pretty.Label("Created: ")
	fmt.Printf("%s%s%s\n", pretty.CyanString(""), aliasData.CreatedAt, pretty.ResetString())
	pretty.Separator()
	return nil
}

func Search(args string) error {
	query := strings.TrimSpace(args)
	if query == "" {
		pretty.Error("No query specified")
		fmt.Println("Usage: tolo search query")
		return fmt.Errorf("no query specified")
	}

	results := storage.SearchAliases(query)

	if len(results) == 0 {
		pretty.Info(fmt.Sprintf("No aliases found matching '%s'", query))
		return nil
	}

	pretty.Header(fmt.Sprintf("Search Results: '%s'", query))

	maxNameLen := 0
	for _, a := range results {
		if len(a.Name) > maxNameLen {
			maxNameLen = len(a.Name)
		}
	}

	for i, a := range results {
		fmt.Printf("  %s%d%s  %s%-*s%s  %s→%s  %s%s%s\n",
			"\033[2m", i+1, "\033[0m",
			"\033[1m", maxNameLen, a.Name, "\033[0m",
			"\033[36m", "\033[0m",
			"\033[32m", a.Command, "\033[0m")
	}

	pretty.Separator()
	fmt.Printf("  Found: ")
	pretty.Count(len(results))
	return nil
}

func Completion(args string) error {
	if args == "" {
		aliases := storage.ListAliases()
		for _, a := range aliases {
			fmt.Println(a.Name)
		}
		return nil
	}

	query := strings.ToLower(args)
	aliases := storage.ListAliases()
	for _, a := range aliases {
		if strings.HasPrefix(strings.ToLower(a.Name), query) {
			fmt.Println(a.Name)
		}
	}
	return nil
}

func GenerateBashCompletion() string {
	return `_tolo_completion() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    if [[ ${COMP_CWORD} -eq 1 ]]; then
        opts="save s run r update u delete del rm d list ls l show sh info search se find help h version v"
        COMPREPLY=($(compgen -W "${opts}" -- "${cur}"))
    elif [[ ${COMP_CWORD} -eq 2 ]]; then
        case "${prev}" in
            run|r|update|u|delete|del|rm|d|show|sh|info|search|se|find)
                COMPREPLY=($(tolo --completion "${cur}"))
                ;;
        esac
    fi
}

complete -F _tolo_completion tolo`
}

func GenerateZshCompletion() string {
	return `#compdef tolo

_tolo() {
    local -a commands
    commands=(
        'save:Save a new alias'
        's:Save a new alias (shortcut)'
        'run:Execute a saved alias'
        'r:Execute a saved alias (shortcut)'
        'update:Update an existing alias'
        'u:Update an existing alias (shortcut)'
        'delete:Delete an alias'
        'del:Delete an alias (shortcut)'
        'rm:Delete an alias (shortcut)'
        'd:Delete an alias (shortcut)'
        'list:List all aliases'
        'ls:List all aliases (shortcut)'
        'l:List all aliases (shortcut)'
        'show:Show details of an alias'
        'sh:Show details of an alias (shortcut)'
        'info:Show details of an alias (shortcut)'
        'search:Search aliases'
        'se:Search aliases (shortcut)'
        'find:Search aliases (shortcut)'
        'help:Show help'
        'h:Show help (shortcut)'
        'version:Show version'
        'v:Show version (shortcut)'
    )

    if [[ CURRENT -eq 2 ]]; then
        _describe 'command' commands
    elif [[ CURRENT -eq 3 ]]; then
        case $words[2] in
            run|r|update|u|delete|del|rm|d|show|sh|info|search|se|find)
                local aliases
                aliases=($(tolo --completion ''))
                _describe 'aliases' aliases
                ;;
        esac
    fi
}

_tolo`
}

func Help() string {
	return `tolo - Simple command alias manager

Usage:
    tolo <command> [arguments]

Commands:
    save (s)     alias:command    Save a new alias
    run (r)      alias            Execute a saved alias
    update (u)   alias:command    Update an existing alias
    delete (d)   alias            Delete an alias
    list (ls, l)                  List all aliases
    show (sh)    alias            Show details of an alias
    search (se)  query            Search aliases
    help (h)                      Show this help message
    version (v)                   Show version

Examples:
    tolo save server1:ssh user@192.168.1.10
    tolo run server1
    tolo update server1:ssh root@192.168.1.10
    tolo list
    tolo ls
    tolo show server1
    tolo search ssh

Installation:
    Install shell completion:
    source <(tolo --bash-completion)    # for bash
    source <(tolo --zsh-completion)     # for zsh`
}

func Version(version string) string {
	return fmt.Sprintf("tolo version %s", version)
}
