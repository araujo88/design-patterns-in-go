# Singleton

The Singleton pattern is a design pattern that restricts the instantiation of a class to a single instance. This is useful when exactly one object is needed to coordinate actions across the system.

The Singleton pattern is used for logging, driver objects, caching, thread pools, and database connections. It is a relatively simple pattern, but it comes with its own set of complications. The main issue is that it introduces global state into an application. For this reason, it's generally not considered a good practice to use singletons in a system that you want to be easy to test and maintain. However, it can be useful in certain scenarios where you need to coordinate logic in a globally accessible way.

# Example

The singleton `WorkerPool` is the focus of this example. It's a structure containing a pool of workers, a pool of jobs and a `WaitGroup`.

```go
type WorkerPool struct {
	workers    []*Worker
	jobQueue   chan Job
	workerPool chan chan Job
	wg         sync.WaitGroup
}
```

Each `Worker` has a job channel to receive jobs and a quit channel for termination. When the worker is started, it continually listens on its job channel for incoming jobs and signals the `WaitGroup` when the job is done. When it receives a job, it prints a message. If it receives a signal on the quit channel, it stops.

```go
type Worker struct {
	ID            int
	WorkerPool    chan chan Job
	JobChannel    chan Job
	quit          chan bool
}

func NewWorker(workerPool chan chan Job, id int) *Worker {
	return &Worker{
		ID:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				fmt.Printf("Worker %d received job %d\n", w.ID, job.ID)
				wg.Done()
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
```

`Job` is a simple structure with just an ID. It represents a task that needs to be done by a worker.

```go
type Job struct {
	ID int
}
```

The singleton `WorkerPool` is initialized using `GetWorkerPool`, which ensures the singleton nature using `sync.Once`. Inside `GetWorkerPool`, a number of workers are started equal to `maxWorkers`. A dispatch function is also started that continually listens for incoming jobs on the jobs channel of the `WorkerPool`, and assigns each job to a worker. The `AddJob` method increases the `WaitGroup` counter and sends the job into the `jobQueue`. The `dispatch` method pulls jobs from `jobQueue` and sends them to available workers, ensuring that jobs are distributed to workers as they become available. The `Wait` method of the `WorkerPool` waits for all jobs to finish. In the main function, we call this method after all jobs have been added to ensure that the program doesn't exit until all jobs have been processed.

```go
func GetWorkerPool(maxWorkers int) *WorkerPool {
	once.Do(func() {
		workers := make([]*Worker, maxWorkers)
		workerPool := make(chan chan Job, maxWorkers)
		jobQueue := make(chan Job, maxWorkers)

		pool := &WorkerPool{
			workers:    workers,
			jobQueue:   jobQueue,
			workerPool: workerPool,
		}

		for i := 0; i < maxWorkers; i++ {
			worker := NewWorker(pool.workerPool, i)
			worker.Start(&pool.wg)
			workers[i] = worker
		}

		go pool.dispatch()

		singleton = pool
	})

	return singleton
}

func (wp *WorkerPool) dispatch() {
	for job := range wp.jobQueue {
		workerJobQueue := <-wp.workerPool
		workerJobQueue <- job
	}
}

func (wp *WorkerPool) AddJob(job Job) {
	wp.wg.Add(1)
	wp.jobQueue <- job
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}
```

This code will start 5 workers and process 20 jobs. Each job will be picked up by an available worker and processed concurrently. The main function will wait for all jobs to be processed before exiting.

```go
func main() {
	wp := GetWorkerPool(5)

	for i := 1; i <= 20; i++ {
		job := Job{ID: i}
		wp.AddJob(job)
	}

	wp.Wait()
}
```
