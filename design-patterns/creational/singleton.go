package creational

import (
	"fmt"
	"sync"
)

/*
# Singleton Pattern

## Concept
This pattern ensures that only one object of its kind exists and provides a
single point of access to it for any other code.

## Example explanation
In this pattern we create a getInstance function to instantiate a certain struct
only once, to ensure that (in case of goroutines) we use sync.Once.
*/

// Singleton
var once sync.Once

type single struct{}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			},
		)
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

// Client Code
func RunSingletonExample() {
	for i := 0; i < 30; i++ {
		go GetInstance()
	}

	fmt.Scanln()
}
