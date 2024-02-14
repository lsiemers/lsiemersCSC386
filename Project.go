package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	for {
		var inputcommand string
		fmt.Printf("Please enter a command:") // Printing Prompt
		if _, err := fmt.Scanln(&inputcommand); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if inputcommand == "exit" { //If user input is the same as the string
			fmt.Println("Exiting shell.") //Print Statement
			return                        //Returning
		}

		switch inputcommand {
		case "ls": //If user input is the same as the string
			out, err := exec.Command("ls").Output()
			if err != nil { //Error Handling
				fmt.Println("Error:", err) //Print Statement
			} else {
				fmt.Println(string(out))
			}

		case "wc": //If user input is the same as the string
			out, err := exec.Command("wc").Output()
			if err != nil { //Error Handling
				fmt.Println("Error:", err) //Print Statement
			} else {
				fmt.Println(string(out))
			}

		case "mkdir": //If user input is the same as the string
			out, err := exec.Command("mkdir").Output()
			if err != nil { //Error Handling
				fmt.Println("Error:", err) //Print Statement
			} else {
				fmt.Println(string(out))
			}

		case "cp": //If user input is the same as the string
			out, err := exec.Command("cp").Output()
			if err != nil { //Error Handling
				fmt.Println("Error:", err) //Print Statement
			} else {
				fmt.Println(string(out))
			}

		case "mv": //If user input is the same as the string
			out, err := exec.Command("mv").Output()
			if err != nil { //Error Handling
				fmt.Println("Error:", err) //Print Statement
			} else {
				fmt.Println(string(out))
			}

		case "whoami": //If user input is the same as the string
			whoami() //Calling whoami Function

		default:
			fmt.Println("Command not supported.")
		}
	}
}

func whoami() {
	// Get user ID Function
	uid := os.Getuid()
	fmt.Printf("User: %s, UserID: %d\n", os.Getenv("USER"), uid)
}
