package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Network struct {
	NumLayer int
	Sizes    []int
	Bias     []float64
	Weight   []float64
}

func NewNetwork(sizes []int) (*Network, error) {
	if len(size) < 2 {
		return nil, fmt.Errorf("must have at least two layer of network")
	}

	Bias := [][]float64{}
	for _, size := range sizes {

	}

	n := &Network{
		NumLayer: len(sizes),
		Sizes:    sizes,
		Bias:     randomArray(len(sizes)),
		Weight:   randomArray(len(sizes)),
	}

	return n, nil
}

func randomArray(size int) []float64 {
	arr := make([]float64, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Float64()
	}
	return arr
}

/*
func (n *Network) feedForward(in []int) {
	return nil
}*/

// Basic Function

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(x))
}

func sigmoidDeri(x float64) float64 {
	return sigmoid(x) * (1 - sigmoid(x))
}
