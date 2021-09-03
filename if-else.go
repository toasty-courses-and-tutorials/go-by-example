package main

import "fmt"

func main() {
    if 7 % 2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8 % 4 == 0 {
        fmt.Println("8 is divisable by 4")
    }

    if num := 9;num < 0 {
        fmt.Println("num is less than 0")
    } else if num < 10 {
        fmt.Println("num has one digit")
    } else {
        fmt.Println("num has more than one digit")
    }

}
