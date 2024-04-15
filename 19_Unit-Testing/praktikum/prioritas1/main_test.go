package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangleArea(t *testing.T) {
	evenTest := RectangleArea(20, 20)
	oddTest := RectangleArea(15, 25)

	assert.Equal(t, "even rectangle", evenTest)
	assert.Equal(t, "odd rectangle", oddTest)
}

func TestRectanglePerimeter(t *testing.T) {
	trueTest := RectanglePerimeter(20, 20)
	falseTest := RectanglePerimeter(15, 27)

	assert.Equal(t, true, trueTest)
	assert.Equal(t, false, falseTest)
}

func TestSquareArea(t *testing.T) {
	evenTest := SquareArea(2)
	oddTest := SquareArea(3)

	assert.Equal(t, "even square", evenTest)
	assert.Equal(t, "odd square", oddTest)
}

func TestSquarePerimeter(t *testing.T) {
	trueTest := SquarePerimeter(10)
	falseTest := SquarePerimeter(8)

	assert.Equal(t, true, trueTest)
	assert.Equal(t, false, falseTest)
}
