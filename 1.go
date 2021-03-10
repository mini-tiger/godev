package main

import (
	"fmt"
	"time"
)

//定义一个实现Job接口的数据
type Score struct {
	Num int
}

//定义对数据的处理
func (s *Score) Do() {
	fmt.Println("num:", s.Num)
	time.Sleep(500 * time.Millisecond) //模拟执行的耗时任务
}

func main() {
	var chan1 = make(chan int, 1)
	var chan2 = make(chan int, 1)
	go func() {
		for {
			select {
			case i := <-chan2:
				fmt.Println(i)
			}
		}

	}()
	go func() {
		chan1 <- 1
		chan2 <- <-chan1
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(len(chan2), len(chan1))

}

// --------------------------- Job ---------------------
type Job interface {
	Do()
}
type JobQueue chan Job

// --------------------------- Worker ---------------------
type Worker struct {
	JobChan JobQueue //每一个worker对象具有JobQueue（队列）属性。
}

func NewWorker() Worker {
	return Worker{JobChan: make(chan Job)}
}

//启动参与程序运行的Go程数量
func (w Worker) Run(wq chan JobQueue) {
	go func() {
		for {
			wq <- w.JobChan //处理任务的Go程队列数量有限，每运行1个，向队列中添加1个，队列剩余数量少1个 (JobChain入队列)
			select {
			case job := <-w.JobChan:
				//fmt.Println("xxx2:",w.JobChan)
				job.Do() //执行操作
			}
		}
	}()
}

// --------------------------- WorkerPool ---------------------
type WorkerPool struct { //线程池：
	Workerlen   int           //线程池的大小
	JobQueue    JobQueue      //Job队列，接收外部的数据
	WorkerQueue chan JobQueue //worker队列：处理任务的Go程队列
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		Workerlen:   workerlen,
		JobQueue:    make(JobQueue),
		WorkerQueue: make(chan JobQueue, workerlen),
	}
}
func (wp *WorkerPool) Run() {
	fmt.Println("初始化worker")
	//初始化worker(多个Go程)
	for i := 0; i < wp.Workerlen; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue) //开启每一个Go程
	}
	// 循环获取可用的worker,往worker中写job
	go func() {
		for {
			select {
			//将JobQueue中的数据存入WorkerQueue
			case job := <-wp.JobQueue: //线程池中有需要待处理的任务(数据来自于请求的任务) :读取JobQueue中的内容
				worker := <-wp.WorkerQueue //队列中有空闲的Go程   ：读取WorkerQueue中的内容,类型为：JobQueue
				worker <- job              //空闲的Go程执行任务  ：整个job入队列（channel） 类型为：传递的参数（Score结构体）
				//fmt.Println("xxx1:",worker)
				//fmt.Printf("====%T  ;  %T======\n",job,worker,)
			}
		}
	}()
}
