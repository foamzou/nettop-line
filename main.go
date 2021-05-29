package main

import (
	"fmt"
	"nettop-line/nettop"
	"os"
	"strings"
	"sync"
	"time"
)
var statChan = make(chan string, 1000)
var outputStrList []string
var statDebounce = debounce(100 * time.Millisecond)

func main () {
	params := strings.Join(os.Args[1:], " ")
	if strings.TrimSpace(params) == "" {
		fmt.Println(
			"Args is required\n" +
				"`nettop -h` to look up usage\n" +
				"For Example: ./nettop-line -P -d -L 0 -J bytes_in,bytes_out -t external -s 1 -c")
		os.Exit(0)
	}
	go nettop.Start(statChan, params)

	for {
		processStr := <- statChan
		outputStrList = append(outputStrList, processStr)
		statDebounce(output)
	}
}

func output()  {
	if len(outputStrList) <= 0 {
		return
	}
	fmt.Println(strings.Join(outputStrList, "|SPLIT|"))

	// clear
	outputStrList = make([]string, 0)
}


func debounce(interval time.Duration) func(f func()) {
	var l sync.Mutex
	var timer *time.Timer

	return func(f func()) {
		l.Lock()
		defer l.Unlock()

		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(interval, f)
	}
}