package main

import (
	"fmt"
	"time"

	"github.com/nitingupta8876/GoTraining/mypackage"
)

func init() {
	fmt.Println("1st Main init")
}

func main() {
	fmt.Println("Hello, Modules!")

	//currTime := time.Now()
	currTime := time.Date(2023, 2, 13, 9, 24, 30, 45, time.Local)

	fmt.Println("The time is ", currTime)

	mypackage.PrintHello()
}

func init() {
	fmt.Println("2nd Main init")
}
