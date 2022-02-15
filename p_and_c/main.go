package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

//生产者 消费者模型
func main() {
	//创建两个通道，一个存放int值，一个存放string值
	ch1 := make(chan int)
	ch2 := make(chan string)

	//开启3个协程
	wg.Add(3)
	go products1(ch1)
	go products2(ch2)
	go customers1(ch1, ch2)
	wg.Wait()
}

//生产者模型1-int
//定时发送
func products1(ch1 chan int) {
	for i := 0; ; i++ {
		ch1 <- i
		time.Sleep(time.Duration(time.Second))
	}
}

//生产者模型1-string
func products2(ch2 chan string) {
	for i := 0; ; i++ {
		ch2 <- strconv.Itoa(i + 5)
		time.Sleep(time.Duration(time.Second))
	}
}

func customers1(ch1 chan int, ch2 chan string) {
	//创建一个定时器发送空数据，避免锁死
	chRate := time.Tick(time.Duration(time.Second * 5))

	//循环消费数据
	for {
		select {
		case value := <-ch1:
			fmt.Printf("接收数据 %T 值 %v \n", value, value)
		case value := <-ch2:
			fmt.Printf("接收数据 %T 值 %v \n", value, value)
		case <-chRate:
			fmt.Printf("Log log...\n")
		}
	}
}
