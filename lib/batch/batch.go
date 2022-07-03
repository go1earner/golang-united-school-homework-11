package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var mx sync.Mutex
	var wg sync.WaitGroup
	wg.Add(int(pool))
	step := n / pool
	for i := int64(0); i < n; i += step {
		go func(low, high int64) {
			for i := low; i < high; i++ {
				result := getOne(i)
				mx.Lock()
				res = append(res, result)
				mx.Unlock()
			}
			wg.Done()
		}(i, i+step)
	}
	wg.Wait()
	return res
}
