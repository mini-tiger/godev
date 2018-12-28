package main

import (
	"errors"
	"fmt"
	"time"
	"log"
)

func RunFunc(t int) (err error) {
	if t != 0 {

		return errors.New(fmt.Sprintf("this is RunFunc ERROR"))
	}
	time.Sleep(time.Duration(100)*time.Hour)
	return nil
}

func TestTimeout() (err error) {
	timeout := time.Duration(5 * time.Second)

	done := make(chan error, 1)
	go func() {
		err := RunFunc(0) // todo 这里显示 运行超时
		//err := RunFunc(1) // todo 这里显示 错误
		done <- err
	}()

	select {

	case err := <-done: // 如果报错 返回
		if err != nil {
			//log.Println(err)
			return err
		}
	case <-time.After(timeout): //如果超时
		err=errors.New(fmt.Sprintf("运行超时"))
		return err
	}


	return
}

func main() {
	err:=TestTimeout()
	log.Println(err)
	select {}
}
