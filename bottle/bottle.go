package main

import (
	"fmt"
	"strconv"
)

const MAX_BOTTLES = 99

func main() {
	for i := MAX_BOTTLES; i >= 0; i-- {
		if i == 0 {
			fmt.Print("No more bottles of beer on the wall, no more bottles of beer.\n")
			fmt.Printf("Go to the store and buy some more, %v bottles of beer.\n\n", MAX_BOTTLES)
		} else {
			fmt.Printf("%v bottles of beer on the wall, %v bottles of beer.\n", i, i)
			var remainAfter string = strconv.Itoa(i - 1)
			if i == 1 {
				remainAfter = "no more"
			}
			fmt.Printf("Take one down and pass it around, %s bottles of beer.\n\n", remainAfter)
		}

	}
}
