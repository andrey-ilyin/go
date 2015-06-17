package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"

    "github.com/andrey-ilyin/fileutil"
)

func main() {
    bads, err := fileutil.ReadLines("bad.txt")
    if err != nil {
        fmt.Printf("Unable to read bad dictionary!")
    }

    goods, err := fileutil.ReadLines("good.txt")
    if err != nil {
        fmt.Printf("Unable to read good dictionary!")
    }

    countWords, countBads, countGoods := 0, 0, 0

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    words := strings.Split(text, " ")
    fmt.Printf("%q\n", words)
    countWords = len(words)

    for _, bad := range bads {
        countBads = countBads + strings.Count(text, bad)
    }

    for _, good := range goods {
        countGoods = countGoods + strings.Count(text, good)
    }

    fmt.Printf("Words: %d bads: %d goods: %d\n", countWords, countBads, countGoods)
}
