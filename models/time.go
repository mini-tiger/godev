package main

import (
	"log"
	"time"
)

//!+main
func bigSlowOperation() int {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(2 * time.Second) // simulate slow operation by sleeping
	return 1
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	log.Println(bigSlowOperation())
}
