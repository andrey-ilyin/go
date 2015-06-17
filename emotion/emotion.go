package main

import "fmt"
import "github.com/andrey-ilyin/fileutil"

func main() {
    bads, err := fileutil.ReadLines("bad.txt")
    if err != nil {
        fmt.Printf("Unable to read bad dictionary!")
    }

    goods, err := fileutil.ReadLines("good.txt")
    if err != nil {
        fmt.Printf("Unable to read good dictionary!")
    }

    for _, bad := range bads {
        fmt.Println(bad)
    }

    for _, good := range goods {
        fmt.Println(good)
    }
}
