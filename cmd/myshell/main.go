package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	hash := map[string]bool{
		"exit": true,
		"echo": true,
		"type": true,
		"pwd":  true,
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSuffix(input, "\n")
		cmd, args := splitByFirstSpace(input)
		switch cmd {
		case "pwd":
			wd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stdout, "Error at getting working directory.")
				break
			}
			fmt.Fprintln(os.Stdout, wd)

		case "type":
			if hash[args] {
				fmt.Fprintln(os.Stdout, fmt.Sprintf("%s is a shell builtin", args))
			} else {
				path, err := exec.LookPath(args)
				if err != nil {
					fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: not found", args))
					break
				}
				fmt.Fprintln(os.Stdout, fmt.Sprintf("%s is %s", args, path))
			}

		case "echo":
			fmt.Fprintln(os.Stdout, args)

		case "exit":
			if args == "0" {
				os.Exit(0)
			} else {
				fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: command not found", cmd))
			}

		default:
			command := exec.Command(cmd, args)
			if output, err := command.Output(); err != nil {
				fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: command not found", cmd))
			} else {
				fmt.Fprintln(os.Stdout, strings.TrimSpace(string(output)))
			}
		}
	}
}

func splitByFirstSpace(s string) (string, string) {
	parts := strings.SplitN(s, " ", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return s, ""
}
