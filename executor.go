package main

import (
	"context"
	"runtime"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

var (
	numCollectors = runtime.NumCPU()

	stop chan struct{}
)

type Handler interface {
	RegisterHandlers(c *colly.Collector, s *Scrapper)
}

type Executor struct {
	ctx       context.Context
	collector *colly.Collector

	mLinks   *sync.Mutex
	linksCh  chan string
	linksSet map[string]struct{}

	mLinksRetry   *sync.Mutex
	linksRetryMap map[string]int
}

func NewExecutor(collector *colly.Collector) *Executor {
	return &Executor{
		ctx:           context.Background(),
		collector:     collector,
		mLinks:        new(sync.Mutex),
		linksCh:       make(chan string, 1_000_000),
		linksSet:      make(map[string]struct{}),
		mLinksRetry:   new(sync.Mutex),
		linksRetryMap: make(map[string]int),
	}
}

func (ex *Executor) Visit(link string) {
	ex.mLinks.Lock()
	if _, ok := ex.linksSet[link]; ok {
		ex.mLinks.Unlock()
		return
	}
	ex.linksSet[link] = struct{}{}
	ex.mLinks.Unlock()
	ex.linksCh <- link
}

func (ex *Executor) retry(link string) {
	ex.mLinksRetry.Lock()
	if v := ex.linksRetryMap[link]; v > 2 {
		ex.mLinksRetry.Unlock()
		return
	}
	ex.linksRetryMap[link] += 1
	ex.mLinksRetry.Unlock()
	time.Sleep(5 * time.Second)
	ex.linksCh <- link
}

func (ex *Executor) start(link string) {
	wg := new(sync.WaitGroup)
	for i := 0; i < numCollectors; i++ {
		wg.Add(1)
		newInstance(i, ex).start(wg)
	}
	go func() {
		wg.Wait()
		stop <- struct{}{}
	}()
	ex.linksCh <- link
	<-stop
}

type Scrapper struct {
	*Executor
	ctx context.Context
	id  int
	c   *colly.Collector
	ch  chan string
	sch chan struct{}
}

func newInstance(id int, ex *Executor) *Scrapper {
	return &Scrapper{
		Executor: ex,
		ctx:      ex.ctx,
		id:       id,
		c:        ex.collector.Clone(),
		ch:       ex.linksCh,
		sch:      make(chan struct{}, 1),
	}
}

func (inst *Scrapper) GetRunnerId() int {
	return inst.id
}

func (inst *Scrapper) start(wg *sync.WaitGroup) {
	inst.setup()
	go inst.run(wg)
}

func (inst *Scrapper) setup() {
	inst.c.OnError(func(r *colly.Response, err error) {
		inst.retry(r.Request.URL.String())
		<-inst.sch
	})

	inst.c.OnResponse(func(r *colly.Response) {
		<-inst.sch
	})
}

func (inst *Scrapper) run(wg *sync.WaitGroup) {
	for {
		var link string
		select {
		case link = <-inst.ch:
		case <-time.After(30 * time.Second):
			wg.Done()
			return
		}

		inst.sch <- struct{}{}
		if err := inst.c.Visit(link); err != nil {
			<-inst.sch
		}
	}
}
