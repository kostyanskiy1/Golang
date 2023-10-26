package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetTime(host string) (time.Time, error) {
	return ntp.Time(host)
}

func main() {
	time, err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection to NTP-server failed: %s", err)
		os.Exit(1)
	}
	fmt.Printf("NTP-server time: %s", time)
}
