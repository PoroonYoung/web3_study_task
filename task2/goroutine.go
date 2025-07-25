package main

import (
	"fmt"
	"time"
)

func printSingle() {
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func printDouble() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	go printSingle()
	go printDouble()
	time.Sleep(1 * time.Second)

	//任务切片
	tasks := []func(){
		func() { time.Sleep(1 * time.Second) }, // 任务1
		func() { time.Sleep(2 * time.Second) }, // 任务2
	}
	taskSchedule(tasks)
}

type taskInfo struct {
	begin time.Time
	end   time.Time
	id    int
	spend time.Duration
}

// 任务调度器
func taskSchedule(slice []func()) {
	taskInfoChan := make(chan taskInfo, len(slice))
	for index, task := range slice {
		go func(idx int, t func()) {
			var info taskInfo
			info.id = idx + 1
			info.begin = time.Now()
			t()
			info.end = time.Now()
			info.spend = info.end.Sub(info.begin)
			taskInfoChan <- info
		}(index, task)
	}
	for i := 0; i < len(slice); i++ {
		info := <-taskInfoChan
		fmt.Printf("任务%d 执行时间: %v\n", info.id, info.spend)
	}
}
