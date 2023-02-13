package mysubpack

import "fmt"

func init() {
	fmt.Println("mysubpack init")
}

func PrintHello() {
	fmt.Println("Hello, Modules! This is mysubpackage speaking!")
}

