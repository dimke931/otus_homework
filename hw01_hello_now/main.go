package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const timeFormat = "2006-01-02 15:04:05 -0700 MST"

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal("Can not get exact time")
	}
	currentTime := time.Now()
	fmt.Println("current time: " + currentTime.Format(timeFormat))
	fmt.Println("exact time: " + exactTime.Format(timeFormat))
}
