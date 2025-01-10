package faults

import (
	"fmt"
	"runtime"
	"time"
)

// TODO cleanup/cpu - add error handling (when does this fail?)
func InjectCPU(vmName string, duration int) error {
	fmt.Printf("CPU Spike for %d seconds\n", duration)

	numCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numCores)

	stopChannel := make(chan struct{})
	for i := 0; i < numCores; i++ {
		go func() {
			for {
				select {
				case <-stopChannel:
					return
				default:
					for i := 0; i < 1000; i++ {
						_ = i * i
					}
				}
			}
		}()
	}
	time.Sleep(time.Duration(duration) * time.Second)
	close(stopChannel)

	fmt.Println("CPU Spike test complete.")
	return nil
}
