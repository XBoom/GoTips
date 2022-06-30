package main

import "gotips/tchan"

func main() {
	ch := make(chan interface{}, 3)
	tchan.Drain(ch)
}
