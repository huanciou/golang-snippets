package main

import (
	"context"
	"sync"
	"time"
)

type MyConcurrentMap struct {
	sync.Mutex
	mp     map[int]int
	key2ch map[int]chan struct{}
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mp:     make(map[int]int),
		key2ch: make(map[int]chan struct{}),
	}
}

func (m *MyConcurrentMap) Put(key int, value int) {
	m.Lock()
	defer m.Unlock()

	m.mp[key] = value

	if ch, exists := m.key2ch[key]; !exists {
		return
	} else {

		// ch <- struct{}{}
		// 這邊這樣做不行，高併發下可能有多個在等待的 goroutine，這樣只會有一個會被喚醒

		// close(ch)
		// 當一個 chan 被關閉，所有讀取阻塞的 goroutine 都會被喚醒

		// 新問題：如果 ch 已經被先前關閉，會怎麼樣？
		// ch 如果關閉兩次，會 Fatal。怎麼解決？

		// <- ch 我直接在讀一次，
		// 這邊如果你讀到東西，代表 ch 已經被關過了
		// 如果你直接阻塞了，就代表 ch 還沒被關閉
		// 但問題是，你驗證了 ch 還沒被關閉，但你也把自己阻塞起來了，會造成死鎖

		select {
		case <-ch:
			return
		default:
			close(ch)
		}

	}
}

func (m *MyConcurrentMap) Get(key int, maxWaitTime time.Duration) (int, error) {
	m.Lock()

	if v, exists := m.mp[key]; exists {
		m.Unlock()
		return v, nil
	}

	// 如果 key 不存在, 需要將 ch 設立成一個阻塞態

	ch, exists := m.key2ch[key]
	if !exists {
		m.key2ch[key] = make(chan struct{})
	}

	tCtx, cancel := context.WithTimeout(context.Background(), maxWaitTime)
	defer cancel()

	m.Unlock()

	select {
	case <-tCtx.Done():
		return -1, tCtx.Err()
	case <-ch:
	}

	m.Lock()
	v := m.mp[key]
	m.Unlock()

	return v, nil
}
