package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

func RunPinger(ctx context.Context, wg *sync.WaitGroup, host string, count int, duration time.Duration, interval time.Duration, timeout time.Duration) {
	defer wg.Done()
	t := time.NewTicker(duration)
	str := fmt.Sprintf("\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "Error", "PacketsSent", "PacketsRecv",
		"PacketLoss", "MinRtt", "AvgRtt", "MaxRtt")
	log.Print(str)
	for {
		select {
		case <-ctx.Done():
			log.Println("shutting down")
			return
		case <-t.C:
			var str string
			pinger, err := ping.NewPinger(host)
			if err != nil {
				str = fmt.Sprintf("%v", err)
			} else {
				pinger.SetPrivileged(true)
				pinger.Count = count
				pinger.Timeout = timeout
				pinger.Interval = interval
				stats, err := doPing(pinger)
				if err != nil {
					str = fmt.Sprintf("%v", err)
				} else {
					str = fmt.Sprintf("%s\t%d\t%d\t%.2f\t%v\t%v\t%v",
						"", stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss,
						stats.MinRtt, stats.AvgRtt, stats.MaxRtt)
				}
			}
			str = fmt.Sprintf("\t%s\n", str)
			log.Print(str)
		}
	}
}

func doPing(pinger *ping.Pinger) (*ping.Statistics, error) {
	var err error
	var stats *ping.Statistics
	err = pinger.Run()
	if err != nil {
		return stats, err
	}
	stats = pinger.Statistics()
	return stats, nil
}
