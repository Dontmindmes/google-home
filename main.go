package main

import (
	"fmt"
	"time"

	"github.com/evalphobia/google-home-client-go/googlehome"
	"github.com/micro/mdns"
)

func main() {
	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()

	cli, err := googlehome.NewClientWithConfig(googlehome.Config{
		Hostname: "192.168.0.2",
		Lang:     "ja",
		Accent:   "GB",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("ai")

	// Change language
	cli.Notify("かねたさんこんにちは、Googleです。")
	// Start the lookup
	close(entriesCh)
	time.Sleep(time.Second * 5)
}
