package main

import (
	"fmt"

	"github.com/antbiz/antapi/cmd/bench/execute"
	"github.com/antbiz/antapi/cmd/bench/sync"
	"github.com/gogf/gf/os/gcmd"
)

var (
	helpContent = `
USAGE
	bench COMMAND [ARGUMENT] [OPTION]
COMMAND
	execute	   execute patch file
	sync	   sync schemas/projects/defaults
OPTION
	-y         all yes for all command without prompt ask 
	-?,-h      show this help or detail for specified command
	-v,-i      show version information
ADDITIONAL
	Use 'bench help COMMAND' or 'bench COMMAND -h' for detail about a command, which has '...' 
	in the tail of their comments.`
)

func help(command string) {
	switch command {
	case "execute":
		execute.Help()
	default:
		fmt.Println(helpContent)
	}
}

func main() {
	command := gcmd.GetArg(1)
	if gcmd.ContainsOpt("h") && command != "" {
		help(command)
		return
	}
	switch command {
	case "help":
		help(gcmd.GetArg(2))
	case "execute":
		execute.Run()
	case "sync":
		sync.Run()
	default:
		fmt.Println(helpContent)
	}
}
