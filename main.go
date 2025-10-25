package main

import (
	"fmt"
	"os"
	"strings"
)

// print short help message
func usage() {
	fmt.Println(`TODO - a tiny todo CLI using GO
Usage:
	todo add "Tasks"
	todo list
	todo done <id>
	todo rm <id>`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		title := strings.Join(os.Args[2:], " ")
		fmt.Println("Adding:", title)
	case "list":
	case "done":
	case "rm":
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Println("Unknown Command:", cmd)
		fmt.Println()
		usage()
		os.Exit(1)
	}
}
