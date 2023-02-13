package mypackage

import (
	"fmt"

	_ "github.com/nitingupta8876/GoTraining/mypackage/mysubpack"
)

func init() {
	fmt.Println("mypackage init")
}

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}
