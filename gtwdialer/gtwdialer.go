// Package gtwdialer pings a series of IP addresses concurrently in the style of old war dialers
package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	maxThreads = 349 // This should be changed to reflect the maximum for your system
)

type IP struct {
	a, b, c, d int
}

func main() {
	ip := IP{192, 168, 0, 0}
	comm := make(chan *string)
	var str *string

	for curThreads := 0; curThreads <= maxThreads; curThreads++ {
		go ip.Ping(comm)
		ip.Next()
	}

	for {
		str = <-comm
		if str != nil {
			fmt.Println(*str)
		}

		for runtime.NumGoroutine() < maxThreads {
			ip.Next()
			go ip.Ping(comm)
		}
	}
}

// Ping attempts to ping a address communicating sucess over chan c
func (ip IP) Ping(c chan *string) {
	// TODO: Replace this with native Go code, impossible at the moment
	cmd := exec.Command("ping", "-c", "2", "-i", "0.2", "-w", "1", ip.ToString())
	exit := cmd.Run()

	if exit == nil {
		str := ip.ToString()
		c <- &str
	}
	c <- nil
}

// Next advances the IP address specified by i
func (i *IP) Next() {
	//Ugly as fuck. Needs to be cleaned
	if i.c%5 == 0 && i.d == 0 {
		fmt.Println("Checkin: ", i.ToString())
	}
	i.d = i.d + 1
	if i.d > 255 {
		i.d = 0
		i.c += 1

		if i.c > 255 {
			i.c = 0
			i.b += 1

			if i.b > 255 {
				i.b = 0
				i.a += 1
			}
		}
	}
}

func (i *IP) ToString() string {
	return fmt.Sprintf("%d.%d.%d.%d", i.a, i.b, i.c, i.d)
}
