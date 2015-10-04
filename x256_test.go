package x256

import (
  "testing" 
  "github.com/stretchr/testify/assert"  
)

func TestClosestCode(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(ClosestCode(220,40,150), 162, "Closest index must be 162")
}