// Package lf is a smaller way to fatally log errors in Go. This package should be 
// considered BAD FOR EVERYONE and not used anywhere. This was mainly an experiment
// for playing with runtime data
package lf

import (
	"runtime"
	"time"
	"strings"
	"os"
	"fmt"
)

func Log(e error) {
	if (e != nil) {
		_, str, line, _ := runtime.Caller(2)
		i:= strings.LastIndex(str, "/")
		t:= time.Now()
		fmt.Printf("[FATAL] [%v:%v] [%vh %vm %vs] [%s]\n", str[i+1:], line, t.Hour(), t.Minute(), t.Second(), e)
		os.Exit(1)
	}
}
