package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	for{
		fmt.Printf("%v\n",os.Getenv("TOMCATHOME"))
		time.Sleep(time.Duration(2)*time.Second)
	}

}