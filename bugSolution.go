```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 5) // Buffered channel

        for i := 0; i < 5; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        val, ok := <-ch
                        if ok {
                                fmt.Printf("Goroutine %d received %d from the channel\n", i, val)
                        } else {
                                fmt.Printf("Goroutine %d: Channel closed\n", i)
                        }
                }(i)
        }

        for i := 0; i < 5; i++ {
                ch <- i
        }
        close(ch)
        wg.Wait()
}
```