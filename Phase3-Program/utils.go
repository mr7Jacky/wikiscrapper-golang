package digit_recognition

import (
	"fmt"
	"math"
	"math/rand"
)

type Network struct {
	numLayer int
	size     []int
	bias     []float64
	weight   []float64
}

func NewNetwork(size []int) (*Network, error) {
	if len(size) < 2 {
		return nil, fmt.Errorf("must have at least two layer of network")
	}

	n := &Network{
		numLayer: len(size),
		size:     size,
		bias:     randomArray(len(size)),
		weight:   randomArray(len(size)),
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

func (n *Network) feedForward(in []int) {
	for _, b := range n.bias {
		for _, w := range n.weight {
			in = mat.dot
		}
	}
}

// Basic Function

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(x))
}

func sigmoidDeri(x float64) float64 {
	return sigmoid(x) * (1 - sigmoid(x))
}
