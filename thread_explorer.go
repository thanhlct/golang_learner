package main
import (
    "fmt"
    "time"
    "sync"
)
func say(s string, id *int) {
    /*
    if *id==30{
        time.Sleep(3*time.Second)
    }
    */
    mux := new(sync.Mutex)
    mux.Lock()
    idInternal := *id
    mux.Unlock()
    time.Sleep(100*time.Millisecond)

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(i, s , idInternal)
    }
}

func main() {
    var i int=30
    id := &i
    go say("------------world", id)
    time.Sleep(100*time.Millisecond)
    i = 20
    go say("=============hello", id)
    *id = 10
    say("*****************thanh", id)
}
