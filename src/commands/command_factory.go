package commands

import (
	"flag"
	"fmt"
	"os"
)

const (
	ListFlag   = "list"
	SetFlag    = "set"
	RemoveFlag = "remove"
)

const (
	CodeFlag        = "c"
	WorkspaceFlag   = "w"
	NameShortFlag   = "n"
	DescriptionFlag = "d"
	VerboseFlag     = "v"
)

func CreateCommand(command string, args []string) Action {
	var action Action
	switch command {
	case SetFlag:
		flagSet := flag.NewFlagSet(SetFlag, flag.ExitOnError)
		workspaceFlagValue := flagSet.Bool(WorkspaceFlag, false, "If project is a workspace or not")
		codeFlagValue := flagSet.String(CodeFlag, "", "Code for the project")
		descriptionFlagValue := flagSet.String(DescriptionFlag, "", "Description of the project")
		parseFlags(flagSet, args)
		checkNumberOfArgs(flagSet)
		validateRequiredFlag(codeFlagValue, CodeFlag)
		if *codeFlagValue == SetFlag || *codeFlagValue == ListFlag || *codeFlagValue == RemoveFlag {
			fmt.Println("This keyword are reserved. Please use another code.")
			os.Exit(0)
		}
		projectPath := flagSet.Arg(0)
		action = SetCommand{Code: *codeFlagValue, Workspace: *workspaceFlagValue, Description: descriptionFlagValue, ProjectPath: projectPath}
	case ListFlag:
		flagSet := flag.NewFlagSet(ListFlag, flag.ExitOnError)
		verboseFlag := flagSet.Bool(VerboseFlag, false, "Shows verbose output")
		parseFlags(flagSet, args)
		action = ListCommand{Verbose: *verboseFlag}
	case RemoveFlag:
		flagSet := flag.NewFlagSet(RemoveFlag, flag.ExitOnError)
		codeFlagValue := flagSet.String(CodeFlag, "", "Given code name of the project")
		parseFlags(flagSet, args)
		validateRequiredFlag(codeFlagValue, CodeFlag)
		action = RemoveCommand{Code: *codeFlagValue}
	default:
		action = CodeCommand{Code: command}
	}
	return action
}

func checkNumberOfArgs(flagSet *flag.FlagSet) {
	if flagSet.NArg() == 0 {
		fmt.Println("Must provide path to a project")
		fmt.Printf("%-8s for the current working directory.\n", ".")
		fmt.Printf("%-8s for a specific path.\n", "<path>")
		os.Exit(1)
	}
}

func validateRequiredFlag(flagValue *string, flag string) {
	if len(*flagValue) == 0 {
		fmt.Printf("the following argument is required: -%v\n", flag)
		os.Exit(0)
	}
}

func parseFlags(flagSet *flag.FlagSet, followingArgs []string) {
	err := flagSet.Parse(followingArgs)
	if err != nil {
		fmt.Println("Invalid parameters.")
		os.Exit(1)
	}
}
