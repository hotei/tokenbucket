// example.go (c) 2014 David Rook - all rights reserved

package main



import (
	// go 1.2 stdlib pkgs
	"fmt"
	"sync"
	"time"
	
	// local pkgs
	"github.com/hotei/tokenbucket"
	)

var hist map[int]int

func main() {
	hist := make(map[int]int)
	tok := tokenbucket.New(time.Millisecond * 1, 1)	// add 1000 tokens per second
	startTime := time.Now()
	iter := 0
	var wg sync.WaitGroup
	var lok sync.Mutex
	for i:=0; i< 29; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				//tok.Dump()
				time.Sleep(tok.Take(1))
				//tok.Wait(5)
				lok.Lock()
				fmt.Printf("gort[%d] %d \n",id, iter)
				hist[id]++
				iter++
				if iter > 1000 { 
					lok.Unlock()
					break 
				}
				lok.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("All done in %v\n",time.Now().Sub(startTime))
	fmt.Printf("Histogram of goroutine id and hits\n")
	fmt.Printf("Fair rate limit would imply nearly equal bucket counts\n")
	for i:=0; i<=50; i++ {
		v,exists := hist[i]
		if exists {
			fmt.Printf("%d has %d hits\n",i,v)
		}
	}
}