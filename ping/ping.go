package main

import (
	"log"
	"os"
	"runtime/trace"
	"time"
	"fmt"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
	for {
		received := <- channel
		fmt.Println("Foo is receiving", received)
		channel <- "ping"
		fmt.Println("Foo is sending: ping")
	}

}

func bar(channel chan string) {
	// TODO: Write an infinite loop of receiving "pings" and sending "pongs"
	for {
		channel <- "pong"
		fmt.Println("Bar is sending: pong")
		received := <- channel
		fmt.Println("Bar is receiving", received)
   }
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	channel := make(chan string, 1)

	go foo(channel) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(channel)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
