package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		command = strings.TrimSpace(command)
		args := strings.Split(command, " ")
		command = args[0];
		
		args = args[1:]
		if command == "exit" {
			if args[0] == "0" {
				return
			}
		}
		if command == "echo" {
			fmt.Println(strings.Join(args, " "))
			continue
		}
		fmt.Printf("%s: command not found\n", command)
	}

}
