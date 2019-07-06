//go test -v hello_test.go -test.run TestHello

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	name := "Bob"
	age := 20

	assert.Equal(t, "Bob", name)
	assert.Equal(t, 20, age)
}

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	object := "Something"
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", object)

	}

}

func TestSomething2(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(123, 456, "they should not be equal")

	object := "Something"
	// assert for nil (good for errors)
	//assert.Nil(object)

	// assert for not nil (good when you expect something)
	if assert.NotNil(object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("Something", object)
	}
}
