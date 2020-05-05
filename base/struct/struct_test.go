package _struct

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

var ageRWLock sync.RWMutex

func (p *Person) older() {
	ageRWLock.Lock()
	defer ageRWLock.Unlock()
	p.Age++
}

func TestStruct(t *testing.T) {
	bob := Person{
		Name:    "Bob",
		Age:     0,
		Address: "xian",
	}
	t.Log(bob)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bob.older()
		}()
	}
	wg.Wait()
	t.Log(bob)
}

func checkType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int:
		fmt.Printf("Inter %v\n", reflect.ValueOf(v))
	case reflect.Bool:
		fmt.Printf("Bool %v\n", reflect.ValueOf(v))
	default:
		fmt.Printf("unknown %v\n", reflect.ValueOf(v))
	}
}

func TestReflect(t *testing.T) {
	checkType(1)
	checkType(false)
	checkType(1.22)
}
