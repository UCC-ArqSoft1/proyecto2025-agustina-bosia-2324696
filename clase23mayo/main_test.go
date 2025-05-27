package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMayor(t *testing.T) {
	result := mayor(5, 10)
	fmt.Println("el resultado es: ", result)
	assert.Equal(t, 10, result)
	assert.Equal(t, 5, mayor(5, 2))
	assert.Equal(t, 0, mayor(0, 0))
}
