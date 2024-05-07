package creational

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	time.Sleep(5 * time.Second)
}

func TestSingletonOnce(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstanceOnce()
	}

	time.Sleep(5 * time.Second)
}

type single struct{}

var lock = &sync.Mutex{}
var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already created.")
		}
	} else {
		fmt.Println("single instance already created.")
	}

	return singleInstance
}

var once sync.Once

func getInstanceOnce() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
