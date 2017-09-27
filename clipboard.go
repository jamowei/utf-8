package main

import (
	"os"
	"strconv"
	"fmt"
	"github.com/atotto/clipboard"
)

func main() {
	var codepoint string
	if len(os.Args) > 1 {
		codepoint = os.Args[1]
	} else {
		fmt.Print("codepoint?: ")
		fmt.Scanln(&codepoint)
	}
	char, err := strconv.Unquote(`"\u` + codepoint + `"`)
	if err != nil {
		fmt.Errorf("codepoint %s is not valid utf-8 char\n", codepoint)
		os.Exit(1)
	}
	if err := clipboard.WriteAll(char); err != nil {
		fmt.Errorf("could not insert to clipboard: %s\n", err)
	} else {
		fmt.Println("utf-8 char copied to clipboard...")
	}
}
