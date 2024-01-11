package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func main() {
	//basics()
	readAndWrite()
}

func readAndWrite() {
	read()
	read()
	write()
	write()
	fmt.Println("Done")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Read unlocking")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlocking")
}

func basics() {
	iterations := 1000

	for i := 0; i < iterations; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is : ", count)
}

func increment() {
	lock.Lock() //lock resource until operation on the resource is done so that no other go routine will use it.
	count++
	lock.Unlock() //unlock to release resource
}
