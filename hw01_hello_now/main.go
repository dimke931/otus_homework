package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
	"github.com/spf13/cast"
)

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		panic("Can not get exact time")
	}
	currentTime := time.Now()
	fmt.Println("current time: " + cast.ToString(currentTime))
	fmt.Println("exact time: " + cast.ToString(exactTime))
}
