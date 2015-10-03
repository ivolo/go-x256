package x256

import (
  "testing" 
  "github.com/stretchr/testify/assert"
  "fmt"
)

func TestCode(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(Code(220,40,150), 162, "Closest index must be 162")
}