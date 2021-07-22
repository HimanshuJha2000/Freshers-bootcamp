package main

import (
	"fmt"
	"sync"
)

func frequency (wg *sync.WaitGroup, s string, mmap map[string]int)  {
	var mutex = &sync.Mutex{}
	defer wg.Done()
	for i:=0; i<len(s); i++ {
		key := string(s[i])
		mutex.Lock()
		mmap[key]++
		mutex.Unlock()
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	var allWords [1000]string

	for i:=0;i<n;i++ {
		var temp string
		fmt.Scanln(&temp)
		allWords[i] = temp
	}

	mmap := make(map[string]int)

	wg := new(sync.WaitGroup)
	wg.Add(n)

	for i:=0;i<n;i++  {
		var temp string
		temp = allWords[i]
		go frequency(wg, temp, mmap)
	}

	wg.Wait()

	for k, v := range mmap {
		fmt.Println(k, "value is", v)
	}

}