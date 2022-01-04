package main

import (
	"fmt"
	"time"
)

func main() {
	// 有goroutine泄漏问题，必须用带buff的channel 或 使用close，因为可能超时后，栈退出后，就没有再有接收channel的操作
	// 而不带buff的channel，发送和接收必须同时发生，因为超时后栈退出了，没有接收的操作，
	// 就会导致ready一致处于发送被阻塞的状态，goroutine没法退出(这种情况很难被发现)
	ready := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second) // 故意设置比超时时间大
		// close(ready)
		ready <- struct{}{}
		fmt.Println("goroutine exit") // 超时后被阻塞，没执行到这步，产生了gorotuine泄漏
	}()

	select {
	case <-ready:
		fmt.Println("ready")
	case <-time.After(2 * time.Second):
		fmt.Println("exit")
	}

	time.Sleep(time.Hour)
}
