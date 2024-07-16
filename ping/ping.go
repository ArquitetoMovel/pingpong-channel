package ping

import (
	"fmt"
	"sync"
	"time"
)

func Send(pingChannel chan<- int, pongChannel <-chan int, exitChannel <-chan bool, wg *sync.WaitGroup) {
	pingData := 1
	for {
		select {
		case pingChannel <- pingData:
			fmt.Println("⚫︎->Send ping ►\t\t", pingData)
			pingData++
		case pongData := <-pongChannel:
			fmt.Println("⚫︎<-Receive pong ◃\t", pongData)
		case <-exitChannel:
      fmt.Println("⊗ <-Receive exit")
      wg.Done()
      return
		}
    time.Sleep(time.Millisecond * 75)
	}
}
