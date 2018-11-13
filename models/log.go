package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)


type Logstruct struct {
	sync.RWMutex
	Log *log.Logger
}

func (l *Logstruct)Info(str string)  {
	l.Log.SetPrefix("[INFO]")
	l.Log.Print(str)
}

func (l *Logstruct)Error(str string)  {
	l.Log.SetPrefix("[ERROR]")
	l.Log.Print(str)
}

var log11 Logstruct

func main() {
	fmt.Println("begin TestLog ...")
	file, err := os.OpenFile("c:\\mysqlsync.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("fail to create test.log file!", err)
	}
	log11.Log = log.New(file, "", log.Ltime|log.Ldate)
	//log.Println("1.Println log with log.LstdFlags ...")
	//logger.Println("1.Println log with log.LstdFlags ...")


	go func() {
		for{
			log11.Info(fmt.Sprintln(1111))
			time.Sleep(time.Duration(1)*time.Second)
		}
	}()
	go func() {
		for{
			log11.Error(fmt.Sprintln(222))
			time.Sleep(time.Duration(1)*time.Second)
		}
	}()

	//logger.SetFlags(0)
	//for   {
	//logger.Println("2.Println log without log.LstdFlags ...")
	//time.Sleep(2*time.Second)
	//}
	//log.Println("2.Println log without log.LstdFlags ...")
	//logger.Println("2.Println log without log.LstdFlags ...")

	//log.Panicln("3.std Panicln log without log.LstdFlags ...")
	//fmt.Println("3 Will this statement be execute ?")
	//logger.Panicln("3.Panicln log without log.LstdFlags ...")

	//log.Println("4.Println log without log.LstdFlags ...")
	//logger.Println("4.Println log without log.LstdFlags ...")

	//log.Fatal("5.std Fatal log without log.LstdFlags ...")
	//fmt.Println("5 Will this statement be execute ?")
	//logger.Fatal("5.Fatal log without log.LstdFlags ...")

	select {}
}
