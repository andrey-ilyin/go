package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const MAX_BOTTLES = 99
	var remainMessage = "%s bottles of beer"
	var remainOnTheWall = "%s bottles of beer on the wall"
	var takeMessage = "Take one down and pass it around"
	var gotoStoreMessage = "Go to the store and buy some more"
	var noMoreMessage = "No more"

	for i := MAX_BOTTLES; i >= 0; i-- {
		if i == 0 {
			fmt.Printf(remainOnTheWall + ", " + remainMessage + ".\n", noMoreMessage, strings.ToLower(noMoreMessage))
			fmt.Printf(gotoStoreMessage + ", " + remainMessage + ".\n\n", strconv.Itoa(MAX_BOTTLES))
		} else {
			fmt.Printf(remainOnTheWall + ", " + remainMessage + ".\n", strconv.Itoa(i), strconv.Itoa(i))
			var remainAfter string = strconv.Itoa(i-1)
			if i == 1 {
				remainAfter = strings.ToLower(noMoreMessage)
			}
			fmt.Printf(takeMessage + ", " + remainMessage + ".\n\n", remainAfter)
		}

	}
}
