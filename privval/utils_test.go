package privval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	cmn "github.com/XuanMaoSecLab/DolphinChain/libs/common"
)

func TestIsConnTimeoutForNonTimeoutErrors(t *testing.T) {
	assert.False(t, IsConnTimeout(cmn.ErrorWrap(ErrDialRetryMax, "max retries exceeded")))
	assert.False(t, IsConnTimeout(fmt.Errorf("completely irrelevant error")))
}
