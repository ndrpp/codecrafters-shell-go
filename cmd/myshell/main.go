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

		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		cmd = strings.TrimSuffix(cmd, "\n")
		switch cmd {
        case "exit 0":
            os.Exit(0)

		default:
			fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: command not found", cmd))
			break
		}
	}
}
