package office

import (
	"reflect"
	"sync"
)

type Pool map[reflect.Type][]any
type Stack map[reflect.Type][]any

var pool Pool
var stack Stack
var locker sync.Locker

func init() {
	pool = make(Pool)
	stack = make(Stack)
	locker = &sync.Mutex{}
}

func Offer[T any](obj *T) {
	locker.Lock()
	pool[reflect.TypeOf(*obj)] = append(pool[reflect.TypeOf(*obj)], obj)
	locker.Unlock()
}

func Take[T any]() *T {
	r := new(T)
	locker.Lock()
	if len(pool[reflect.TypeOf(*r)]) > 0 {
		obj := pool[reflect.TypeOf(*r)]
		pool[reflect.TypeOf(*r)] = obj[1:]
		locker.Unlock()
		return obj[0].(*T)
	}
	locker.Unlock()
	return r
}

func Push[T any](obj *T) {
	locker.Lock()
	stack[reflect.TypeOf(*obj)] = append(stack[reflect.TypeOf(*obj)], obj)
	locker.Unlock()
}

func Pop[T any]() *T {
	r := new(T)
	locker.Lock()
	if len(stack[reflect.TypeOf(*r)]) > 0 {
		obj := stack[reflect.TypeOf(*r)]
		stack[reflect.TypeOf(*r)] = obj[1:]
		locker.Unlock()
		return obj[0].(*T)
	}
	locker.Unlock()
	return nil
}
