package main

import (
	"fmt"
	"os"

	"github.com/livebud/color"
)

func main() {
	fmt.Fprintln(os.Stdout, color.Red("h"),
		color.Blue("e"),
		color.Yellow("l"),
		color.Green("l"),
		color.Pink("o"),
	)
}
