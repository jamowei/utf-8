package main

import (
	"os"
	"strconv"
	"fmt"
	"github.com/atotto/clipboard"
	"unicode/utf8"
	"strings"
	"encoding/hex"
)

func main() {
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		exception(1, "provide one of the following input types:\n" +
			"  codepoint    [U+XXXX]\n" +
			"  hex value    [XXXX]\n" +
			"  utf-8 symbol [X]\n")
	}
	do(input)
}

func do(input string) {
	if utf8.RuneCountInString(input) > 1 {
		var char string
		var err error
		input = strings.ToLower(input)

		if strings.HasPrefix(input, "u+") {
			// codepoint
			char, err = strconv.Unquote(`"\u` + strings.TrimPrefix(input, "u+") + `"`)
			if err != nil {
				exception(1, "input %s is not valid utf-8 codepoint\n", input)
			}
			if !utf8.ValidString(char) {
				exception(1, "input %s is not valid utf-8 codepoint\n", input)
			}
		} else {
			// hex value
			var res []byte
			res, err = hex.DecodeString(input)
			if err != nil {
				exception(1, "input %s is not valid hex utf-8 value\n", input)
			} else {
				char = string(res)
			}
		}
		copy2clip(char)
	} else {
		// utf-8 char
		if utf8.ValidString(input) {
			codepoint, _ := utf8.DecodeRuneInString(input)
			fmt.Printf("utf8 codepoint of %s is U+%x\n", input, codepoint)
			fmt.Printf("utf8 hex value of %s is %x", input, input)
		}
	}
}

func copy2clip(text string) {
	if err := clipboard.WriteAll(text); err != nil {
		exception(-1, "could not copy to clipboard: %s\n", err)
	} else {
		fmt.Printf("%s copied to clipboard\n", text)
	}
}

func exception(code int, msg string, vars ...interface{}) {
	if len(vars) > 0 {
		fmt.Fprintf(os.Stderr, msg, vars)
	} else {
		fmt.Fprint(os.Stderr, msg)
	}
	if code > -1 {
		os.Exit(code)
	}
}
