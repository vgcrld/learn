package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// Create a anon function (ff) and spin it off with go statements.
	for i, v := range []int{5, 10} {
		wg.Add(1)
		ff := func(v, i int) {
			dur := time.Duration(v) * time.Second
			start_msg := fmt.Sprintf("start i=%v, v=%v, dur=%v", i, v, dur)
			end_msg := fmt.Sprintf("end i=%v, v=%v, dur=%v", i, v, dur)
			fmt.Println(start_msg)
			time.Sleep(dur)
			fmt.Println(end_msg)
			wg.Done()
		}
		go ff(v, i)
	}

	wg.Wait()

}
