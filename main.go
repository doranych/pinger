package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var host, output string
	var count int
	var timeout, interval, duration time.Duration

	flag.StringVar(&host, "h", "google.com", "sets host that will be pinged. Might be IP address")
	flag.StringVar(&output, "o", "stdout", "sets file to output log. Might be relative or absolute")
	flag.IntVar(&count, "c", 5, "sets pings per try count")
	flag.DurationVar(&duration, "d", 5*time.Second, "sets ping interval")
	flag.DurationVar(&timeout, "t", 250*time.Millisecond, "sets pings timeout in ms")
	flag.DurationVar(&interval, "i", 50*time.Millisecond, "sets pings interval in ms")
	flag.Parse()

	fmt.Printf("host: %s\noutput: %s\ncount: %d\nduration: %v\ntimeout: %v\ninterval: %v\n", host, output, count, duration, timeout, interval)

	setLogOutput(output)

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ListenInterrupt(wg, cancel)
	go RunPinger(ctx, wg, host, count, duration, interval, timeout)

	wg.Wait()
	log.Println("done")
}

func ListenInterrupt(wg *sync.WaitGroup, cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		defer wg.Done()
		for range c {
			cancel()
			return
		}
	}()
}

func setLogOutput(output string) {
	var wr io.Writer
	var err error
	if output == "stdout" {
		wr = os.Stdout
	} else {
		wr, err = os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.SetOutput(wr)
}
