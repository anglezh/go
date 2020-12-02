package main

import (
	"LoyZhang/gointro/cmd/pipeline"
	"fmt"
)

func main() {
	p := pipeline.InMemSort(pipeline.ArraySource(3,2,6,4,7))
	for {
		num, ok := <-p
		if !ok {
			break
		}
		fmt.Println(num)
	}
}
