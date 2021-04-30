package main

import (
	"flag"
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

	var action commands.Action
	baseCommand := os.Args[1]
	followingArgs := os.Args[2:]

	switch baseCommand {
	case SetCommand:
		flagSet := flag.NewFlagSet(SetCommand, flag.ExitOnError)
		workspaceFlagValue := flagSet.Bool(WorkspaceFlag, false, "If project is a workspace or not")
		codeFlagValue := flagSet.String(CodeFlag, "", "Code for the project")
		descriptionFlagValue := flagSet.String(DescriptionFlag, "", "Description of the project")
		parseFlags(flagSet, followingArgs)
		if flagSet.NArg() == 0 {
			fmt.Println("Must provide path to a project")
			fmt.Printf("%-8s for the current working directory.\n", ".")
			fmt.Printf("%-8s for a specific path.\n", "<path>")
			os.Exit(1)
		}
		if len(*codeFlagValue) == 0 {
			fmt.Printf("the following argument is required: -%v\n", CodeFlag)
			os.Exit(0)
		}
		projectPath := flagSet.Arg(0)
		action = commands.SetCommand{Code: *codeFlagValue, Workspace: *workspaceFlagValue, Description: descriptionFlagValue, ProjectPath: projectPath}
	case ListCommand:
		flagSet := flag.NewFlagSet(ListCommand, flag.ExitOnError)
		verboseFlag := flagSet.Bool(VerboseFlag, false, "Shows verbose output")
		parseFlags(flagSet, followingArgs)
		action = commands.ListCommand{Verbose: *verboseFlag}
	case RemoveCommand:
		flagSet := flag.NewFlagSet(RemoveCommand, flag.ExitOnError)
		codeFlagValue := flagSet.String(CodeFlag, "", "Given code name of the project")
		parseFlags(flagSet, followingArgs)
		if len(*codeFlagValue) == 0 {
			fmt.Printf("the following argument is required: -%v\n", CodeFlag)
			os.Exit(0)
		}
		action = commands.RemoveCommand{Code: *codeFlagValue}
	default:
		action = commands.CodeCommand{Code: baseCommand}
	}
	action.Execute()
	os.Exit(0)
}

func parseFlags(flagSet *flag.FlagSet, followingArgs []string) {
	err := flagSet.Parse(followingArgs)
	if err != nil {
		fmt.Println("Invalid parameters.")
		os.Exit(1)
	}
}
