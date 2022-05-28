package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	chChromeDie := make(chan struct{})
	chBackendDie := make(chan struct{})
	chSignal := listenToInterrupt()
	//go Run()
	go startBrowser(chChromeDie, chBackendDie)
	for {
		select {
		case <-chSignal:
			chBackendDie <- struct{}{}
			fmt.Println("backend die")
		case <-chChromeDie:
			fmt.Println("received chrome die")
			os.Exit(0)
		}
	}
	for {
		time.Sleep(10 * time.Second)
	}
}
func startBrowser(chChromeDie chan struct{}, chBackendDie chan struct{}) {
	cmd := exec.Command(`open`, `http://127.0.0.1:8888`)
	cmd.Start()
	go func() {
		<-chBackendDie
		cmd.Process.Kill()
	}()
	go func() {
		cmd.Wait()
		chChromeDie <- struct{}{}
	}()
}

func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal)
	return chSignal
}
