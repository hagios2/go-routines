package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	go throwingNinjaStar(channel)
	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	rounds := 3

	for i := 0; i < rounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}

	close(channel)
}
