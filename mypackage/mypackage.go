package mypackage

import (
	"fmt"

	_ "GoTraining/mypackage/mysubpack"
)

func init() {
	fmt.Println("mypackage init")
}

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}
