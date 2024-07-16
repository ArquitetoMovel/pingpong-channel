package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/arquitetomovel/pingpongchannel/ping"
	"github.com/arquitetomovel/pingpongchannel/pong"
)

func main() {
  var w_group sync.WaitGroup
  time_seconds := 2
  w_group.Add(3)
  ping_channel := make(chan int)
  pong_channel := make(chan int)
  exit_ping_channel := make(chan bool)
  exit_pong_channel := make(chan bool)
  fmt.Println(">> Iniciando ping pong por", time_seconds, "segundos <<")

  go ping.Send(ping_channel, pong_channel, exit_ping_channel, &w_group)
  go pong.Send(pong_channel, ping_channel, exit_pong_channel, &w_group) 
 
  go func() {
    fmt.Println("Wait...")
    time.Sleep(time.Second * time.Duration(time_seconds)) 
    exit_pong_channel<- true
    exit_ping_channel<- true
    fmt.Println("Finalizando...")
    w_group.Done()
  }()

  w_group.Wait()
  fmt.Println("Close channels...")

  close(ping_channel)
  close(pong_channel)
  close(exit_ping_channel)
  close(exit_pong_channel)

}
