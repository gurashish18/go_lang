package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time class")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println("Formatted time -> ", presentTime.Format("01-02-2006 Monday 15:04:05")) // it is fixed

	createdDate := time.Date(2020, time.March, 21, 11, 11, 11, 11, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}