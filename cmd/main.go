package main

import (
	"fmt"

	"github.com/TrashWithNoLid/FyteColor/pkg/color"
)

func main() {
	color.SetColor(color.Blue)
	newStr, err := color.BuildStr("Hello World!")
	if err != nil {
		fmt.Printf("Error building string: %v\n", err)
		return
	}

	fmt.Println(newStr)
}
