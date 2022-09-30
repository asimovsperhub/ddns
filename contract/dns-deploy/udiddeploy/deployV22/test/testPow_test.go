package test

import (
	"fmt"
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	n := math.Pow(10, 6)

	fmt.Println(n)
}
