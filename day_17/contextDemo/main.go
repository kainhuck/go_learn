package main

// pakage ==========================================
import (
	"context"
	"fmt"
	"sync"
	"time"
)

// variable ========================================
var (
	wg       sync.WaitGroup
	exitFlag = false
	exitChan chan struct{}
)

// func ============================================

func demoByVariable() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("tutu")
			if exitFlag {
				break
			}
		}
	}()
	time.Sleep(5 * time.Second)
	exitFlag = true
	wg.Wait()
}

func demoByChannel() {
	exitChan = make(chan struct{}, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
	LOOP:
		for {
			fmt.Println("tutu")
			time.Sleep(500 * time.Millisecond)
			select {
			case <-exitChan:
				break LOOP
			default:
			}
		}
	}()
	time.Sleep(5 * time.Second)
	exitChan <- struct{}{}
	wg.Wait()
}

func demoByContext() {
	ctx, cancle := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
	LOOP:
		for {
			fmt.Println("tutu")
			time.Sleep(500 * time.Millisecond)
			select {
			case <-ctx.Done():
				break LOOP
			default:
			}
		}
	}(ctx)
	time.Sleep(5 * time.Second)
	cancle()
	wg.Wait()
}

// main func =======================================
func main() {
	// demoByVariable()
	// demoByChannel()
	demoByContext()
}
