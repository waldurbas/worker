package worker

// ----------------------------------------------------------------------------------
// worker.go package
// Copyright 2019,2023 by Waldemar Urbas
//-----------------------------------------------------------------------------------
// This Source Code Form is subject to the terms of the 'MIT License'
// A short and simple permissive license with conditions only requiring
// preservation of copyright and license notices. Licensed works, modifications,
// and larger works may be distributed under different terms and without source code.
// ----------------------------------------------------------------------------------

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type JobIntf interface {
	Work()
}

type Job struct {
	Worker JobIntf
}

func NewJob() *Job {
	return &Job{}
}

type Service struct {
	maxJobs   int
	startTime time.Time
	Count     uint64
	Finished  uint64
	ch        chan Job
	wg        *sync.WaitGroup
}

func New(maxJobs int) *Service {
	if maxJobs < 1 {
		maxJobs = runtime.GOMAXPROCS(0)
	}

	return &Service{
		maxJobs: maxJobs,
		wg:      &sync.WaitGroup{},
		ch:      make(chan Job, 2),
	}
}

func (s *Service) Start() {
	s.startTime = time.Now()

	for n := 0; n < s.maxJobs; n++ {
		go func() {
			s.workFunc(s.ch)
		}()
	}
}

func (s *Service) workFunc(ch chan Job) {
	for fn := range ch {
		func() {
			defer s.wg.Done()
			fn.Worker.Work()
			atomic.AddUint64(&s.Finished, 1)
		}()
	}
	close(ch)
}

func (s *Service) Add(job Job) {
	s.wg.Add(1)
	s.ch <- job
	atomic.AddUint64(&s.Count, 1)
}

func (s *Service) Wait() {
	s.wg.Wait()
}
