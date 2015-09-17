package main

import (
    "time"
    "fmt"
)
const (
    timeFormat = "2006-01-02 15:04 MST"
)

func sleep(secs int) {
    select {
          case <-time.After(time.Duration(secs) * time.Second):
            return
    }
}

func main() {

	fmt.Println(time.Now())
    sleep(3)
    fmt.Println(time.Now())
}