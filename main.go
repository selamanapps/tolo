package main

import (
	"fmt"
	"os"
	"strings"
	"tolo/cmd"
)

var version = "dev"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(cmd.Help())
		os.Exit(1)
	}

	command := os.Args[1]
	args := strings.Join(os.Args[2:], " ")

	switch command {
	case "save", "s":
		if err := cmd.Save(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "run", "r":
		if err := cmd.Run(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "update", "u":
		if err := cmd.Update(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "delete", "del", "rm", "d":
		if err := cmd.Delete(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "list", "ls", "l":
		if err := cmd.List(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "show", "sh", "info":
		if err := cmd.Show(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "search", "se", "find":
		if err := cmd.Search(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "help", "h":
		fmt.Println(cmd.Help())
	case "version", "v":
		fmt.Println(cmd.Version(version))
	case "--bash-completion":
		fmt.Println(cmd.GenerateBashCompletion())
	case "--zsh-completion":
		fmt.Println(cmd.GenerateZshCompletion())
	case "--completion":
		if err := cmd.Completion(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		fmt.Println(cmd.Help())
		os.Exit(1)
	}
}
