package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	noOfStudents int = 200
	totalScore int
)

func averageScore (wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	totalScore += rand.Intn(100)
	m.Unlock()
	time.Sleep(time.Duration(rand.Float32()))
}


func main() {
	wg := new(sync.WaitGroup)
	wg.Add(200)

	var m sync.Mutex

	for i:=1; i<=200; i++ {
		go averageScore(wg, &m)
	}
	wg.Wait()

	fmt.Println((float32(totalScore/200.0)))
}