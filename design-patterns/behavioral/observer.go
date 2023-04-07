package behavioral

import "fmt"

/*
# Observer Pattern

## Concept
Allows some objects to notify other objects about changes in their state

## Example explanation

*/

// Subject
type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

// Concrete subject
type Task struct {
	observerList []Observer
	name         string
	status       bool
}

func newTask(name string) *Task {
	return &Task{
		name: name,
	}
}

func (t *Task) updateStatus() {
	fmt.Printf("%s status has changed\n", t.name)
	t.status = true
	t.notifyAll()
}

func (t *Task) register(o Observer) {
	t.observerList = append(t.observerList, o)
}

func (t *Task) deregister(o Observer) {
	t.observerList = removeFromSlice(t.observerList, o)
}

func (t *Task) notifyAll() {
	for _, observer := range t.observerList {
		observer.update(t.name)
	}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// Observer
type Observer interface {
	update(string)
	getID() string
}

// Concrete observer
type Manager struct {
	id string
}

func (m *Manager) update(taskName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", m.id, taskName)
}

func (m *Manager) getID() string {
	return m.id
}

// Client code
func RunObserverExample() {
	task1 := newTask("Task 1")

	observerFirst := &Manager{id: "test@mail.com"}
	observerSecond := &Manager{id: "secondtest@mail.com"}

	task1.register(observerFirst)
	task1.register(observerSecond)

	task1.updateStatus()
}
