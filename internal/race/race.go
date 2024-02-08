package race

import (
	"fmt"
	"sync"
)

func Run() {
	var data int
	var wg sync.WaitGroup

	// 複数のゴルーチンで同じ変数を更新
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			data++ // 共有データへのアクセス（レースコンディションの原因）
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Data:", data)
}
