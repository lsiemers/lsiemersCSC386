package main

//
//	Lukas Siemers - Project#1 - CSC386
//
import (
	"fmt"     //Imports
	"os"      //''
	"os/exec" //''
)

func main() {
	var inputcommand string
	fmt.Printf("Please enter a command:") // Printing Prompt
	fmt.Scan(inputcommand)                // Taking User-Input
	for inputcommand != "exit" {
		if inputcommand == "ls" { //If user input is the same as the string
			out, err := exec.Command("ls").Output()
			output := string(out[:])
			fmt.Println(output)
			if err != nil { //Error Handeling (This took me a while to figure out, otherwise the .Run command was not gonna work)
				fmt.Println("Error:", err) //Print Statement
			}

		} else if inputcommand == "wc" { //If user input is the same as the string
			out, err := exec.Command("wc").Output()
			output := string(out[:])
			fmt.Println(output)
			if err != nil { //Error Handeling (This took me a while to figure out, otherwise the .Run command was not gonna work)
				fmt.Println("Error:", err) //Print Statement
			}
		} else if inputcommand == "mkdir" { //If user input is the same as the string
			out, err := exec.Command("mkdir").Output()
			output := string(out[:])
			fmt.Println(output)
			if err != nil { //Error Handeling (This took me a while to figure out, otherwise the .Run command was not gonna work)
				fmt.Println("Error:", err) //Print Statement
			}
		} else if inputcommand == "cp" { //If user input is the same as the string
			out, err := exec.Command("cp").Output()
			output := string(out[:])
			fmt.Println(output)
			if err != nil { //Error Handeling (This took me a while to figure out, otherwise the .Run command was not gonna work)
				fmt.Println("Error:", err) //Print Statement
			}
		} else if inputcommand == "mv" { //If user input is the same as the string
			out, err := exec.Command("mv").Output()
			output := string(out[:])
			fmt.Println(output)
			if err != nil { //Error Handeling (This took me a while to figure out, otherwise the .Run command was not gonna work)
				fmt.Println("Error:", err) //Print Statement
			}
		} else if inputcommand == "whoami" { //If user input is the same as the string
			whoami() //Calling whoami Function
		} else if inputcommand == "exit" { //If user input is the same as the string
			fmt.Println("Exiting shell.") //Print Statement
			return                        //Returning
		} else {
			fmt.Println("Command not supported.")
		}
	}
}

func whoami() {
	// Get user ID Function
	uid := os.Getuid() //I could not comeup on my own with thism so I found this line in the internet
	fmt.Printf("User: %s, UserID: %d\n", os.Getenv("USER"), uid)
}
