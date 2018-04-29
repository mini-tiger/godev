package main

//https://studygolang.com/articles/9532
import (
	"fmt"
	// "time"
)

var complete chan int = make(chan int)      //无缓冲通道
var bufchannel chan int = make(chan int, 6) //缓冲通道

func loop() {
	for i := 0; i < 10; i++ {
		// fmt.Printf("%d ", i)
		complete <- i
	}
	// fmt.Println()
	// complete <- 0 // 执行完毕了，发个消息
}
func nobufferchanel() {

	go loop()
	/*go*/ func() { //匿名函数,如果两个都是go 则并行，上面存通道，这边取出来
		for i := 0; i < 9; i++ {
			// x := <-complete
			fmt.Printf("%d ", <-complete)
		}
	}()
	// time.Sleep(1000 * time.Millisecond) // 1秒
	// 直到线程跑完, 取到消息. main在此阻塞住
}
func main() {
	nobufferchanel()
	fmt.Printf("%d \n", <-complete)

	bufferchanel()
}

func bufferchanel() {
	go func() {
		for i := 0; i < 6; i++ { //不能超过6
			bufchannel <- i * i
		}
		defer close(bufchannel)
	}()

	func() {
		for x := range bufchannel { //range时 语法上立面已经close通道，运行时可以并行
			fmt.Printf("%d ", x)

		}
		fmt.Println()
	}()
	fmt.Println(len(bufchannel))

	// func() {
	// 	for {
	// 		if x, ok := <-bufchannel; ok { //判断，当 不知道 通道是否有多少数据
	// 			fmt.Printf("%d", x)
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }()
	// fmt.Println()

}
