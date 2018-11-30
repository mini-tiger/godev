package test

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
	"godev/test_测试/business"
	"testing"
)

func tt1()  {
	var wg sync.WaitGroup
	wg.Add(200)

	for range [100]struct{}{} {
		go func() {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)

			//log.Println(business.Data())
			wg.Done()
		}()
	}

	for i := range [100]struct{}{} {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)
			s := fmt.Sprint("#", i)
			//log.Println("====", s)

			business.SetData(s)
			wg.Done()
		}(i)
	}

	wg.Wait()

	//fmt.Println("final data = ", *business.data)
	return
}

func tt2()  {
	var wg sync.WaitGroup
	wg.Add(200)
	var ss business.Ldata
	for range [100]struct{}{} {
		go func() {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)

			//log.Println(ss.Data())
			wg.Done()
		}()
	}

	for i := range [100]struct{}{} {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)
			s := fmt.Sprint("#", i)
			//log.Println("====", s)

			ss.SetData(s)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func Benchmark_Tt1(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		tt1()
	}
}

func Benchmark_Tt2(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		tt2()
	}
}