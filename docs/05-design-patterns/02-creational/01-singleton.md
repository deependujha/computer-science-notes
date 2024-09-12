# Singleton Design Pattern

- Only one instance of a class will be available at any time.
- E.g., a single instance of a database connection pool.

```go
// main.go file
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
```

- If we'd used something like:

```go
func getInstance(wg *sync.WaitGroup) *single {
    defer wg.Done()

    // check if instance is already created, to avoid latency bcoz of locks
    lock.Lock()
    defer lock.Unlock()
    if singleInstance == nil {
        fmt.Println("Creating single instance now.")
        singleInstance = &single{}
    } else {
        fmt.Println("Single instance already created.")
    }
    return singleInstance
}
```

- `locking` would have been a bottleneck. So, better to first make sure that the instance is not created, and only then lock and create it.

---

## But, why use `lock` for singleton?

- If different threads are trying to create the single instance, then possibly both of them will be able.
- So, we need to use `lock` to make sure that only one thread is able to create the single instance.
