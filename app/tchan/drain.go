package tchan

import "fmt"

func Drain(channel <- chan interface{}) {
	for range channel {
		fmt.Print("hello world")
	}
}

