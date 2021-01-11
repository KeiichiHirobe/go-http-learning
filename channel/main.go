package main

import (
	"fmt"
	"sync"
	"time"
)

const debug = false

type responseAndError struct {
	response int
	err      error
}

type adder struct {
	// for cancel or timeout
	close chan struct{}
	resc  chan responseAndError
}

func newAdder() *adder {
	return &adder{
		close: make(chan struct{}),
		// must use unbuffered channel
		resc: make(chan responseAndError),
		// if use buffered channel, we would get FAILED log
		// resc:  make(chan responseAndError, 1),
	}
}

func (adder *adder) handle(a int, b int) bool {
	adder.add(a, b)
	time.Sleep(time.Second * 1)
	select {
	case <-adder.close:
		return false
	case rc := <-adder.resc:
		if debug {
			fmt.Printf("result: %d, err: %v", rc.response, rc.err)
		}
		return true
	}

}

func (adder *adder) add(a int, b int) {
	go func(a int, b int) {
		defer func() {
			close(adder.close)
		}()
		res := a + b
		adder.resc <- responseAndError{res, nil}
	}(a, b)
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			adder := newAdder()
			ok := adder.handle(i, 1)
			if !ok {
				fmt.Println("================FAILED============")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
