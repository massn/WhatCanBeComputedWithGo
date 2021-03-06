package chap2

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestYes(t *testing.T) {
	result := ContainsGAGA("CTGAGAC")
	assert.Equal(t, "yes", result)
}

func TestNo(t *testing.T) {
	result := ContainsGAGA("CTGA")
	assert.Equal(t, "no", result)
}

func TestFile(t *testing.T) {
	bytes, err := ioutil.ReadFile("geneticString.txt")
	if err != nil {
		panic(err)
	}
	result := ContainsGAGA(string(bytes))
	assert.Equal(t, "yes", result)
}
