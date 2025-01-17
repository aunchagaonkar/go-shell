package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
		cmd := args[0]
		args = args[1:]
		switch cmd {
		case "exit":{
				if len(args) == 0 {
					fmt.Printf("%s: command not found\n", command)
					continue
				}
				args, err := strconv.Atoi(args[0])
				if err != nil {
					os.Exit(1)
				}
				os.Exit(args)
			}
		case "type":{
				if len(args) == 0 {
					fmt.Printf("%s: this command needs atleast 1 argument\n", cmd)
					continue
				}
				switch args[0] {
				case "exit", "echo", "type", "pwd", "cd":{
						fmt.Printf("%s is a shell builtin\n", args[0])
						continue
					}
				default:{
						f := false
						paths := strings.Split(os.Getenv("PATH"), ":")
						for _, path := range paths {
							fp := filepath.Join(path, args[0])
							if _, err := os.Stat(fp); err == nil {
								fmt.Println(fp)
								f = true
							}
						}
						if !f {
							fmt.Printf("%s: not found\n", args[0])
							continue
						}
					}
				}
			}
		case "echo":{
				fmt.Println(strings.Join(args, " "))
				continue
			}
		case "pwd":{
				dir, err := os.Getwd()
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(dir)
				continue;
		}
		case "cd":{
				if len(args) == 0 || args[0] == "~" {
					os.Chdir(os.Getenv("HOME"))
					continue
				}
				err := os.Chdir(args[0])
				if err != nil {
					fmt.Printf("cd: %s: No such file or directory\n", args[0])
					continue
				}
				continue
		}
		default:{
				// fmt.Printf("%s: command not found\n", cmd)
				cmd := exec.Command(cmd, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Printf("%s: command not found\n", cmd)
				}
				continue
			}
		}
	}
}
