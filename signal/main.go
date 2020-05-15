package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

}

func signal1(){
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT)

	for v := range sig {
		fmt.Println(v)
		signal.Stop(sig)
		close(sig)
	}
}

func signal2() {
	sigRecv1 := make(chan os.Signal, 1)
	sig1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Println("those sig2:",sig1)
	signal.Notify(sigRecv1, sig1...)

	sigRecv2 := make(chan os.Signal, 1)
	sig2 := []os.Signal{syscall.SIGQUIT}
	fmt.Println("those sig2:",sig2)
	signal.Notify(sigRecv2, sig2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Println("sig1:", sig)
		}
		fmt.Println("End. [sigRecv1]")
		wg.Done()
	}()

	go func() {
		for sig := range sigRecv2 {
			fmt.Println("sig2:", sig)
		}
		fmt.Println("End. [sigRecv2]")
		wg.Done()
	}()

	fmt.Println("waiting 2 seconds")
	time.Sleep(2 *time.Second)
	fmt.Println("stop notification...")
	signal.Stop(sigRecv1)
	close(sigRecv1)

	wg.Wait()
}