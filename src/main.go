package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/seanervinson/vc/commands"
)

// Turn this into a map
const SetCommand = "set"
const CodeFlag = "c"
const WorkspaceFlag = "w"
const NameShortFlag = "n"
const DescriptionFlag = "d"
const ListCommand = "list"
const VerboseFlag = "v"

const ConfigCommand = "config"
const PathFlag = "path"

const HelpCommand = "help"
const RemoveCommand = "remove"

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Custom help %s:\n", os.Args[0])

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(os.Stderr, "    %v\n", f.Usage)
	})
}

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
	switch os.Args[1] {
	case SetCommand:
		flagSet := flag.NewFlagSet(SetCommand, flag.ExitOnError)
		workspaceFlagValue := flagSet.Bool(WorkspaceFlag, false, "If project is a workspace or not")
		codeFlagValue := flagSet.String(CodeFlag, "", "Code for the project")
		descriptionFlagValue := flagSet.String(DescriptionFlag, "", "Description of the project")
		err := flagSet.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Invalid parameters.")
			os.Exit(1)
		}
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
		err := flagSet.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Invalid parameters.")
			os.Exit(1)
		}
		action = commands.ListCommand{Verbose: *verboseFlag}

	// case ConfigCommand:
	// 	flagSet := flag.NewFlagSet(ConfigCommand, flag.ExitOnError)
	// 	flagSet.String(PathFlag, "", "Update the project path.")
	// 	flagSet.String(DescriptionFlag, "", "Update the project description.")
	case RemoveCommand:
		flagSet := flag.NewFlagSet(RemoveCommand, flag.ExitOnError)
		codeFlagValue := flagSet.String(CodeFlag, "", "Given code name of the project")
		flagSet.Parse(os.Args[2:])
		if len(*codeFlagValue) == 0 {
			fmt.Printf("the following argument is required: -%v\n", CodeFlag)
			os.Exit(0)
		}
		action = commands.RemoveCommand{Code: *codeFlagValue}
	default:
		// Read config.json
		cmd := exec.Command("code")
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Something unexpected happened.")
		}
		os.Exit(0)
	}
	action.Execute()
}
