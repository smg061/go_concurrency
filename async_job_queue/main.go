package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Job interface {
	Id() string;
	Process()
}
type Worker struct {
	done *sync.WaitGroup
	quit chan bool
	readyPool chan chan Job
	assignedJobQueue chan Job
}


type JobQueue struct {
	internalQueue 		chan Job
	readyPool 			chan chan Job
	workers 			[]*Worker
	dispatcherStopped 	*sync.WaitGroup
	workersStopped 		*sync.WaitGroup
	quit 				chan bool
}

func NewWorker(readyPool chan chan Job, done *sync.WaitGroup) *Worker {
	return &Worker{
	  done:             done,
	  readyPool:        readyPool,
	  assignedJobQueue: make(chan Job),
	  quit:             make(chan bool),
	}
  }

func NewJobQueue(maxWorkers int) *JobQueue {
	workersStopped := sync.WaitGroup{}
	readyPool := make(chan chan Job, maxWorkers)
	workers := make([]*Worker, maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		workers[i] = NewWorker(readyPool, &workersStopped)
	}

	return &JobQueue{
		internalQueue: make(chan Job),
		readyPool: readyPool,
		workers: workers,
		dispatcherStopped: &sync.WaitGroup{},
		workersStopped: &workersStopped,
		quit: make(chan bool),
	}
}

func (q *JobQueue) Start() {

	for i:= 0; i< len(q.workers); i++ {
		q.workers[i].Start()
	}

	go q.Dispatch()
}

func (q *JobQueue) Stop() {
	q.quit <- true
	q.dispatcherStopped.Wait()
	close(q.internalQueue)
	close(q.readyPool)
}

func (q *JobQueue) Dispatch() {
	q.dispatcherStopped.Add(1)
	for {
		select {
		case job := <- q.internalQueue:
			workerChannel := <-q.readyPool
			workerChannel <- job
		case <- q.quit:
			for i := 0; i <len(q.workers); i++ {
				q.workers[i].Stop()
			}
			q.workersStopped.Wait()
			q.dispatcherStopped.Done()
			return
		}
	}
}

func (q *JobQueue) Submit(job Job) {
	q.internalQueue <- job
}

func (w *Worker) Start() {
	w.done.Add(1)
	go func() {
		for {
			w.readyPool <- w.assignedJobQueue // check the job queue in
			select {
			case job := <- w.assignedJobQueue:
				job.Process()
				fmt.Printf("")
			case <- w.quit:
				w.done.Done()
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.quit <- true
}

type TestJob struct {
	ID string
}
func (t *TestJob) Process() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("Processing job %s \n", t.ID)
	duration := time.Duration(rand.Intn(8))
	time.Sleep(duration * time.Second)
	fmt.Printf("Job %s completed \n", t.ID)
}

func (t *TestJob) Id () string {
	return t.ID
}
func main () {

	queue:= NewJobQueue(runtime.NumCPU())

	queue.Start()
	fmt.Printf("%d\n",runtime.NumCPU())

	defer queue.Stop()

	for i := 0; i < 4*runtime.NumCPU(); i++ {
		queue.Submit(&TestJob{strconv.Itoa(i)})
	}
}