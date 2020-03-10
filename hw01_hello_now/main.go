package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
	"github.com/prometheus/common/log"
)

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal("Can not get exact time")
	}
	currentTime := time.Now()
	fmt.Println("current time: " + currentTime.Format("15:04:05.00000"))
	fmt.Println("exact time: " + exactTime.Format("15:04:05.00000"))
}
