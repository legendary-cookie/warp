package main

import "fmt"

func info(msg string) {
	fmt.Println("[INFO]", msg)
}


func warn(msg string) {
	fmt.Println("[WARN]", msg)
}

func logerr(msg string) {
	fmt.Println("[ERR]", msg)
}
