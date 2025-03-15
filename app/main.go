package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	for {

		fmt.Fprint(os.Stdout, "$ ")                                // Show prompt
		command, err := bufio.NewReader(os.Stdin).ReadString('\n') // Read input

		if err != nil {
			log.Fatal(err)
		}

		if command == "" { //If the user presses Enter without typing anything, the loop continues and prompts again
			continue
		}

		fmt.Println(strings.TrimSpace(command) + ": command not found") // Remove extra spaces
	}

}
