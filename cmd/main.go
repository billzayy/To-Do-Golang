package main

import (
	"github.com/billzayy/to-do-list/pkg"
	"github.com/billzayy/to-do-list/ui"
)

func main() {
	result := &pkg.Todo{}

	for {
		options := ui.Header()

		if options != 5 {
			ui.Content(options, result)
		} else {
			return
		}

		result.ToPrint()
	}
}
