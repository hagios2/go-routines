package main

import "fmt"

func main() {
    //specify channel capacity
    channel := make(chan string, 2)

    channel <- "First message"
    channel <- "Second message"

    //FIFO queue
    fmt.Println(<-channel)
    fmt.Println(<-channel)
}
