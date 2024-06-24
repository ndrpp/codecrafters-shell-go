package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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
		case "echo":
			fmt.Fprintln(os.Stdout, args)

		case "exit":
			if args == "0" {
				os.Exit(0)
			} else {
				fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: command not found", cmd))
			}

		default:
			fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: command not found", cmd))
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
