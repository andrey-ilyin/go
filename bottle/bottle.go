package main

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX_BOTTLES = 99

func main() {
	var noMoreMessage = "No more"

	for i := MAX_BOTTLES; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%s bottles of beer on the wall, %s bottles of beer.\n", noMoreMessage, strings.ToLower(noMoreMessage))
			fmt.Printf("Go to the store and buy some more, %v bottles of beer.\n\n", MAX_BOTTLES)
		} else {
			fmt.Printf("%v bottles of beer on the wall, %v bottles of beer.\n", i, i)
			var remainAfter string = strconv.Itoa(i-1)
			if i == 1 {
				remainAfter = strings.ToLower(noMoreMessage)
			}
			fmt.Printf("Take one down and pass it around, %s bottles of beer.\n\n", remainAfter)
		}

	}
}
