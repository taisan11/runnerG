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
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}

		switch input {
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		case "c":
			fmt.Println("Exiting...")
			os.Exit(0)
		case "clear":
			os.Stdout.WriteString("\033[H\033[2J")
		default:
			args := strings.Split(input, " ")
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				log.Println("Error executing command:", err)
			}
		}
	}
}
