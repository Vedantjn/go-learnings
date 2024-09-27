package main

import "time"

func main() {
	go greeter("Hello") // goroutine -> lightweight thread managed by Go runtime
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		println(s)
	}
}
