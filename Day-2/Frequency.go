package main

import (
	"fmt"
	"sort"
	"sync"
)

func frequency (wg *sync.WaitGroup, s string, mmap map[string]int, m *sync.Mutex)  {

	defer wg.Done()
	for i:=0; i<len(s); i++ {
		key := string(s[i])
		m.Lock()
		mmap[key]++
		m.Unlock()
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
	var m sync.Mutex
	for i:=0;i<n;i++  {
		var temp string
		temp = allWords[i]
		go frequency(wg, temp, mmap, &m)
	}

	wg.Wait()

	keys := make([]string, 0, len(mmap))
	for k := range mmap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, "value is", mmap[k])
	}

}