package MultiThread

import "sync"

func LoopPrint100() {
	var mutex sync.Mutex
	waitGroup := sync.WaitGroup{}
	count := 1
	max := 100
	waitGroup.Add(3)
	go func() {
		defer waitGroup.Done()
		for {
			mutex.Lock()
			if count > max {
				mutex.Unlock()
				break
			}
			if count%3 == 0 {
				println(count)
				count++
			}
			mutex.Unlock()
		}
	}()
	go func() {
		defer waitGroup.Done()
		for {
			mutex.Lock()
			if count > max {
				mutex.Unlock()
				break
			}
			if count%3 == 1 {
				println(count)
				count++
			}
			mutex.Unlock()
		}
	}()
	go func() {
		defer waitGroup.Done()
		for {
			mutex.Lock()
			if count > max {
				mutex.Unlock()
				break
			}
			if count%3 == 2 {
				println(count)
				count++
			}
			mutex.Unlock()
		}
	}()
	waitGroup.Wait()
}
