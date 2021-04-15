package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_greet(t *testing.T) {
	got := greet()

	assert.Equal(t, "Hi, this is Leo!", got, "should properly greet")
}
