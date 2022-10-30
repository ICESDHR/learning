package main

import (
	"sync"
)

type LeakBucket struct {
	m           sync.Mutex
	rate        int
	lastTime    int
	sleepTime   int
	maxInternal int
}

var NowTime = 123

func getNextTime(l LeakBucket) int {
	l.m.Lock()
	defer l.m.Unlock()

	if l.lastTime <= 0 {
		return NowTime
	}

	l.sleepTime += NowTime - l.lastTime

	if l.sleepTime > 1/l.rate {
		return NowTime
	}

	if l.sleepTime > l.maxInternal {
		l.sleepTime = l.maxInternal
	}

	return l.sleepTime
}
