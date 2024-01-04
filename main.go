package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func()  {
		fmt.Println(time.Since(start))
	}()

	ninjas := []string{"Derrick", "Prince", "Max", "jonny"}

	for _, ninja := range ninjas {
		go attack(ninja)
	}

	time.Sleep(time.Second * 2)
}


func attack(target string) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(time.Second)
}
