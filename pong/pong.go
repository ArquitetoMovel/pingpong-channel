package pong

import (
	"fmt"
	"sync"
  "time"
)

func Send(pongChannel chan<- int, pingChannel <-chan int, exitChannel <-chan bool, wg *sync.WaitGroup) {
	pongData := 1
	for {
		select {
		case pongChannel <- pongData:
			fmt.Println("⚪︎->Send pong ►\t\t", pongData)
			pongData++
		case pingData := <-pingChannel:
			fmt.Println("⚪︎<-Receive ping ◃\t", pingData)
		case <-exitChannel:
      fmt.Println("⊗ <-Receive exit")
      wg.Done()
      return
		}
    time.Sleep(time.Millisecond * 100)
	}
}
