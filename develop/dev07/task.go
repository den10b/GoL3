package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		go func(c <-chan interface{}) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		//sig(2*time.Hour),
		//sig(5*time.Minute),
		sig(1*time.Second),
		//sig(1*time.Hour),
		//sig(1*time.Minute),
		sig(10*time.Second),
	)
	fmt.Printf("fone after %v", time.Since(start))
}
