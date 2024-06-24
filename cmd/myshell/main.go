package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	cmd = strings.TrimSuffix(cmd, "\n")
	switch cmd {
	default:
		fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: command not found", cmd))
	}
}
