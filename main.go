package main

import (
	"fmt"
	"gotips/app/mapreduce"
)

func main() {
	v, err := mapreduce.Run([]int{1,2,0,3,4,5})
	fmt.Printf("result %v, err %v", v, err)
}
