package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// source: https://gosamples.dev/unix-time/

	date := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(date.Unix())
	fmt.Println(date.UnixMilli())
	fmt.Println(date.UnixMicro())
	fmt.Println(date.UnixNano())

	// Get current Unix/Epoch time in seconds
	fmt.Println(time.Now().Unix())

	// Print in UTC
	t := time.Unix(time.Now().Unix(), 0)
	fmt.Println(t.UTC())

	// Parse timestamp from string
	timeString := "1757861095"
	timestamp, err := strconv.ParseInt(timeString, 10, 64)
	if err != nil {
		panic(err)
	}
	t = time.Unix(timestamp, 0)
	fmt.Println(t.UTC())

}
