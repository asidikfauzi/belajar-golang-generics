package belajar_golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	result := Length[string]("Fauzi")
	assert.Equal(t, "Fauzi", result)

	resultNumber := Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
