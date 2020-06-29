package Queue

import "sync"

type CQueue struct {
	mem  []int
	lock sync.Mutex
}

func Constructor() CQueue {
	return CQueue{
		mem:  make([]int, 0),
		lock: sync.Mutex{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.mem = append(this.mem, value)
}

func (this *CQueue) DeleteHead() int {
	this.lock.Lock()
	defer this.lock.Unlock()
	if len(this.mem) > 0 {
		last := this.mem[0]
		this.mem = this.mem[1:]
		return last
	}
	return -1
}
