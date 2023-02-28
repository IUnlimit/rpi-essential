package gpio

import "testing"
import "github.com/stretchr/testify/assert"

func TestInit(t *testing.T) {
	pin, err := Builder().SetMode(IN).Init(1).Build()
	assert.NoError(t, err)
	err = pin.SetLevel(UP)
	assert.NoError(t, err)
	err = pin.Close()
	assert.NoError(t, err)
}
