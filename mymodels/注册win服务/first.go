// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// simple does nothing except block while running the service.
package main

import (
	"log"

	"fmt"
	"github.com/kardianos/service"
	"os"
	"time"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() { // todo 可以使用exec.command() 加入写过的程序
	file, err := os.OpenFile("d:\\mysqlsync.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("fail to create test.log file!", err)
	}
	logger1 := log.New(file, "", log.Ltime|log.Ldate)
	//log.Println("1.Println log with log.LstdFlags ...")
	logger1.Println("1.Println log with log.LstdFlags ...")

	for {
		logger1.Println("2.Println log without log.LstdFlags ...")
		time.Sleep(2 * time.Second)
	}
	//log.Println("2.Println log without log.LstdFlags ...")

}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func (p *program) Install(s service.Service) error {
	err := s.Install()
	if err != nil {
		fmt.Printf("Failed to install: %s\n", err)
		return err
	}
	fmt.Printf("Service \"%s\" installed.\n", "install system server")
	return nil
}

func (p *program) Uninstall(s service.Service) error {
	err := s.Uninstall()
	if err != nil {
		fmt.Printf("Failed to install: %s\n", err)
		return err
	}
	fmt.Printf("Service \"%s\" installed.\n", "install system server")
	return nil
}

var s service.Service

func main() {
	svcConfig := &service.Config{
		Name:        "GoServiceExampleSimple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Printf("service.new: %s\n", err)
	}
	//s.Uninstall()
	//s.Install()
	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			err := s.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
			}
		case "remove":
			err = s.Uninstall()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" removed.\n", "go")
		case "run":
			s.Run()
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" started.\n", "go")
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" stopped.\n", "go")
		}
		return
	}

	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
