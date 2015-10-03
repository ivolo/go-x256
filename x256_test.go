package x256

import (
  "testing" 
  "github.com/stretchr/testify/assert"
)

func TestClosestIndex(t *testing.T) {
  assert := assert.New(t)
  code := Code(220,40,150)
  assert.Equal(code, 162, "Closest index must be 162")
}