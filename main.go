package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	phonebook := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Phonebook")
	fmt.Println("Available commands: add,get,update,exit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		// exemple : add ~ alice=123
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
	}
}
