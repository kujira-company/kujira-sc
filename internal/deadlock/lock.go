package deadlock

import "sync"

func Run() {
	var lock1, lock2 sync.Mutex

	lock1.Lock()
	lock2.Lock()

	go func() {
		lock1.Lock()
		lock2.Unlock()
	}()

	go func() {
		lock2.Lock()
		lock1.Unlock()
	}()

	// ここでデッドロック発生
}
