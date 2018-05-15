package main

import (
	"gopkg.in/abiosoft/ishell.v2"
)

// This file should be able to execute the module - this is just the interface
func Run_CLI() {
	shell := ishell.New()
	shell.Print("Welcome to Counzl"
	shell.Print("Type hjelp to get a list of commands :)")
	shell.Println()
	// register a function for "greet" command.

	// 
	shell.AddCmd(&ishell.Cmd{
		Name: "test", // name of the command
		Help: "", // description of the command; this will be displayed if you type 'hjelp'
		Func: func(c *ishell.Context) {
			
			fmt.Println("This is the function body for test-command")

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "another command", // the name should not be as long as this
		Help: "this is another command",
		Func: func(c *ishell.Context) {
			
			fmt.Println("This is the function body to 'another command'")
			
		},
	})

	shell.Run()
}
