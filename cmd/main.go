package main

import (
	"context"
	"fmt"
	"log"
)

func Horse(name string, ch chan<- string, ctx context.Context, cancel context.CancelFunc, count int) {
	i := 0
outer:
	for {
		if i == count {
			log.Printf("%s : Reach the finish line！", name)
			cancel()
			ch <- name
		}
		select {
		default:
			log.Println(name, fmt.Sprintf(": %d m", i+1))
			i++
		case <-ctx.Done():
			break outer
		}
	}
}

func main() {
	ch := make(chan string)
	c := context.Background()
	ctx, cancel := context.WithCancel(c)

	for i := 0; i < 4; i++ {
		go Horse(fmt.Sprintf("Horse-%d", i+1), ch, ctx, cancel, 100)
	}
	winner := <-ch
	log.Printf("%s : Win", winner)
}
