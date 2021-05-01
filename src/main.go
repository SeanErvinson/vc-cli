package main

import (
	"fmt"
	"os"

	"github.com/seanervinson/vc/commands"
)

const (
	CodeFlag        = "c"
	WorkspaceFlag   = "w"
	NameShortFlag   = "n"
	DescriptionFlag = "d"
	VerboseFlag     = "v"
)
const (
	ListCommand   = "list"
	SetCommand    = "set"
	RemoveCommand = "remove"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("vs is a tool to open VSCode project via codenames.")
		fmt.Println("")
		fmt.Println("vc <code>")
		fmt.Println("Opens Visual Studio code base on given code.")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("")
		fmt.Println("\tvc <command> [arguments]")
		fmt.Println("")
		fmt.Println("The commands are: ")
		fmt.Println("")
		fmt.Printf("\t%-8v%-24v\n", "set", "sets a code name base on path")
		fmt.Printf("\t%-8v%-24v\n", "list", "list codes")
		fmt.Printf("\t%-8v%-24v\n", "config", "update code path or description")
		fmt.Printf("\t%-8v%-24v\n", "remove", "remove code")
		os.Exit(0)
	}
	baseCommand := os.Args[1]
	args := os.Args[2:]

	command := commands.CreateCommand(baseCommand, args)
	command.Execute()
}
