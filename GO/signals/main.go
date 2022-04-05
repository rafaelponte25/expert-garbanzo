package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	go channelTest()
	<-killSignal
	fmt.Println("Thanks for using Golang!")
}

func channelTest() {
	for {
		fmt.Println("Doing Work")
		time.Sleep(1 * time.Second)
	}
}
