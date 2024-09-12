package main

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance(wg *sync.WaitGroup) *single {
    defer wg.Done()

    // check if instance is already created, to avoid latency bcoz of locks
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating single instance now.")
            singleInstance = &single{}
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }
    return singleInstance
}

func main() {

    var wg sync.WaitGroup

    for i := 0; i < 30; i++ {
        wg.Add(1)
        go getInstance(&wg)
    }

    wg.Wait()
}
