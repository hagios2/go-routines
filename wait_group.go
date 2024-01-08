package main

import (
	"fmt"
	"sync"
)

func main() {
	var beeper sync.WaitGroup
	beeper.Add(1)
	go attackNinja("Jonny", &beeper)
	beeper.Wait()

	fmt.Println("Mission completed")
}

func attackNinja(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja: ", evilNinja)
	beeper.Done()
}
