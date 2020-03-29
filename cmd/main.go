package main

import (
	"fmt"

	fytecolor "github.com/TrashWithNoLid/FyteColor/pkg/color"
)

func main() {
	err := fytecolor.ColoredPrintf("$$Hello $$World!\n", fytecolor.BrightRed, fytecolor.BrightBlue)
	if err != nil {
		fmt.Printf("Error call ColoredPrintf: %v\n", err)
	}
}
