package goroutinepool

/*
 goroutine pool
 version： 0.1
*/

type Job func()

type worker struct {
	stop chan struct{}
}

func (w *worker) start(jobs chan Job) {
	for {
		select {
		case job := <-jobs:
			job()
		case <-w.stop:
			return
		}
	}
}

type Pool struct {
	worker_size int
	job_queue   chan Job
}

func New(w_size int, j_size int) *Pool {
	return &Pool{
		worker_size: w_size,
		job_queue:   make(chan Job, j_size),
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.worker_size; i++ {
		var worker = &worker{}
		go worker.start(p.job_queue)
	}
}

func (p *Pool) AddJob(job Job) {
	p.job_queue <- job
}
