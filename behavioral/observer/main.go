package main

import "fmt"

// JobPost is a struct that represents a job posting
type JobPost struct {
	title string
}

// Observer interface
type Observer interface {
	Notify(jobPost JobPost)
}

// JobSeeker implements Observer
type JobSeeker struct {
	name string
}

func (js *JobSeeker) Notify(jobPost JobPost) {
	fmt.Printf("Hi %s! New job posted: %s\n", js.name, jobPost.title)
}

// JobBoard is our subject struct
type JobBoard struct {
	subscribers []Observer
}

func (jb *JobBoard) Subscribe(observer Observer) {
	jb.subscribers = append(jb.subscribers, observer)
}

func (jb *JobBoard) AddJob(jobPost JobPost) {
	for _, subscriber := range jb.subscribers {
		subscriber.Notify(jobPost)
	}
}

func main() {
	// Create subscribers
	jobSeeker1 := &JobSeeker{name: "John Doe"}
	jobSeeker2 := &JobSeeker{name: "Jane Doe"}

	// Create publisher and add subscribers
	jobBoard := &JobBoard{}
	jobBoard.Subscribe(jobSeeker1)
	jobBoard.Subscribe(jobSeeker2)

	// Add a new job and see if subscribers get notified
	jobBoard.AddJob(JobPost{"Software Engineer"})
}
