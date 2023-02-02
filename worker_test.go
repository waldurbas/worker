package worker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/waldurbas/worker"
)

type TJob struct {
	id int
}

func (job *TJob) Work() {
	fmt.Println("start:", time.Now().Format("2006-01-02 15:05:05.000"), ", ID:", job.id)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("stop :", time.Now().Format("2006-01-02 15:05:05.000"), ", ID:", job.id)
}

func TestWorker(t *testing.T) {
	startTime := time.Now().Format("2006-01-02 15:05:05.000")
	w := worker.New(0)
	w.Start()

	for a := 0; a < 100; a++ {
		job := worker.Job{
			Worker: &TJob{
				id: a,
			},
		}

		w.Add(job)
	}

	w.Wait()

	fmt.Println("StartTime:", startTime)
	fmt.Println("Count:", w.Count)
	fmt.Println("Finished:", w.Finished)
}
