package main

import "sync"

type Sema struct {
	wg sync.WaitGroup
	ch chan func()
}

func NewSema(n int) *Sema {
	s := &Sema{
		ch: make(chan func(), n),
	}
	for ; n > 0; n-- {
		go s.handler()
	}
	return s
}

func (s *Sema) handler() {
	for fn := range s.ch {
		fn()
		s.wg.Done()
	}
}

func (s *Sema) Run(fn func()) {
	s.wg.Add(1)
	s.ch <- fn
}

func (s *Sema) WaitAndClose() {
	s.wg.Wait()
	close(s.ch)
	s.ch = nil
}
