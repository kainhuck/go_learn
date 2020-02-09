package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second) // 你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。 这里的定时器将等待 2 秒。

	<-timer1.C // <-timer1.C 会一直阻塞， 直到定时器的通道 C 明确的发送了定时器失效的值。
	fmt.Println("Timer1 fired")

	/* 定时器与time.Sleep的不同之处在于定时器可以取消 */
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("定时器关闭")
	}

	time.Sleep(2 * time.Second)
}
