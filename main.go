package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctxLogic(ctx, time.Second)

}

func ctxLogic(ctx context.Context, t time.Duration) {
	select {
	case <-time.After(t):
		mainLogic()
		fmt.Println("works done")
	case <-ctx.Done():
		fmt.Println("Timed out")
	}
}

func mainLogic() {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				res := doWork()
				dataChan <- res

			}()
		}
		wg.Wait()
		close(dataChan)
	}()
	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}
}
