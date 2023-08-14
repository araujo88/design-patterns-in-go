# Observer

The Observer Design Pattern is a behavioral design pattern that allows objects to notify other objects about changes in their state. It establishes a one-to-many dependency between a subject object that holds state, and any number of observer objects that are interested in changes to the subject's state.

Here's a simple analogy: think of a news publisher and its subscribers. When the publisher releases a new article, it notifies all its subscribers. In this scenario, the publisher is the "subject" and the subscribers are the "observers". This concept of automatically notifying interested parties is the core idea behind the Observer pattern.

In software, the Observer pattern is commonly used in event-driven systems. You would use the Observer pattern when changes to the state of one object may require changes in other objects, and the actual number of objects is unknown or dynamic.

Here's a high-level description of the Observer pattern's participants:

- Subject: This is the object that holds the state. It allows observers to register and unregister their interest.

- Observer: These are the objects that want to know about changes. They register their interest with the Subject and also define a method for receiving notifications.

- ConcreteSubject: A concrete subject stores the state of interest to ConcreteObserver objects. When a change occurs in the subject's state, it sends a notification to all registered Observers.

- ConcreteObserver: A ConcreteObserver maintains a reference to a ConcreteSubject object, stores state that should stay consistent with the subject's, and implements the Observer updating interface to keep its state consistent with the subject's.

In Go, this pattern could be implemented using interfaces to define the `Subject` and `Observer`, and individual types that implement these interfaces. The Observer's update method would be called when the `Subject` changes.

## Example

Let's consider a scenario where we have a Job Posting site and users can subscribe to be notified when a new job is posted. This is a good example for the Observer pattern as it involves a subject (Job Postings) and observers (Subscribed Users) where observers need to be notified when there's a change in the subject.

```go
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
```

When a new job is posted (i.e., the JobBoard's state changes), all the subscribed job seekers are notified. This is the Observer pattern in action.
