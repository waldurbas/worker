# worker: 2023.02.02


## Example
```go
type Job struct {
	id   int
	name string
}

func (job *Job) Work() {
	fmt.Println("start.job:", time.Now().Format("2006-01-02 15:05:05"), ", ID:", job.id, ", Name:", job.name)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("stop.job :", time.Now().Format("2006-01-02 15:05:05"), ", ID:", job.id, ", Name:", job.name)
}

func main() {

	w := wrk.New(0)
	w.Start()

	for a := 0; a < 1000; a++ {
		job := wrk.Job{
			Worker: &Job{
				id:   a,
				name: fmt.Sprintf("msg_%d", a),
			},
		}

		w.Add(job)
	}

	w.Wait()
	fmt.Println("Count:", w.Count)
	fmt.Println("Finished:", w.Finished)
}
