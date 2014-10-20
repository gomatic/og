package main

import (
	"os"
)

func SplitArgs(args []string) (string, []string) {
	if len(args) < 2 {
		return "", nil
	}
	return args[1], args[2:]
}

func main() {
	cmd, args := SplitArgs(os.Args)
	og := NewOg(args, ".")
	if cmd == "" {
		cmd = "help"
	}

	switch cmd {
	case "build":
		og.Build()
	case "help":
		og.Help()
	default:
		og.Default(cmd)
	}
}
