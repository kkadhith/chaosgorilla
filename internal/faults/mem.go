package faults

import (
	"fmt"
	"runtime"
	"time"
)

func InjectMem(vmName string, duration int) error {
	fmt.Printf("Memory leak spike for %d seconds\n", duration)

	var memLeaks [][]byte

	stopChannel := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopChannel:
				return
			default:
				kb := 1024
				mb := 1024 * kb

				memChunk := make([]byte, 100*mb)

				for i := 0; i < len(memChunk); i += kb {
					memChunk[i] = 1
				}

				memLeaks = append(memLeaks, memChunk)

				fmt.Printf("Memory leakage at roughly %dMB\n", len(memLeaks)*100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(time.Duration(duration) * time.Second)
	close(stopChannel)

	memLeaks = nil
	runtime.GC()
	fmt.Println("Memory leak Spike test complete.")
	return nil
}
