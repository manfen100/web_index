package main

import "sync"

type Cond struct {
	L sync.Locker
}

func NewCond(l sync.Locker) {
}
