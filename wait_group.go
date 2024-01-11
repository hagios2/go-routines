package main

import (
	"fmt"
	"sync"
)

func main() {
	var beeper sync.WaitGroup
	ninjas := []string{"Jonny", "Marcus", "Prince", "Kevin"}
	beeper.Add(len(ninjas)) //number passed here should be the same as the number of times we call the waitGroup.Done() method
	for _, ninja := range ninjas {
		go attackNinja(ninja, &beeper)
	}

	beeper.Wait()

	fmt.Println("Mission completed")
}

func attackNinja(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja: ", evilNinja)
	beeper.Done()
}
