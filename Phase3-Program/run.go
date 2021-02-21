package main

import (
	"fmt"
)

func main() {
	s := []int{2, 2, 3}
	nw, err := NewNetwork(s)
	if err == nil {
		fmt.Errorf("fail to create network")
	}
	fmt.Println(nw.Bias)

}
